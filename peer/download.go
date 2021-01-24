package peer

import (
    "encoding/binary"
    "math"
    "fmt"

    "github.com/kylec725/graytorrent/common"
    "github.com/kylec725/graytorrent/peer/message"
    "github.com/kylec725/graytorrent/write"
    "github.com/pkg/errors"
)

const maxReqSize = 16384

// Errors
var (
    ErrBitfield = errors.New("Malformed bitfield received")
    ErrMessage = errors.New("Malformed message received")
    ErrPieceHash = errors.New("Received piece with bad hash")
    ErrUnexpectedPiece = errors.New("Received piece when not expecting it")
)

type workPiece struct {
    index int
    piece []byte
    left int  // bytes remaining in piece
}

func (peer *Peer) addWorkPiece(index int) {
    pieceSize := common.PieceSize(peer.info, index)
    piece := make([]byte, pieceSize)
    newWork := workPiece{index, piece, pieceSize}
    peer.workQueue = append(peer.workQueue, newWork)

    // TODO send out work requests, can maybe be a goroutine
}

// getMessage reads in a message from the peer
func (peer *Peer) getMessage() (*message.Message, error) {
    buf := make([]byte, 4)
    if err := peer.Conn.Read(buf); err != nil {
        return nil, errors.Wrap(err, "getMessage")
    }
    msgLen := binary.BigEndian.Uint32(buf)
    if msgLen == 0 {  // Keep-alive message
        return nil, nil
    }

    buf = make([]byte, msgLen)
    err := peer.Conn.ReadFull(buf)
    return message.Decode(buf), errors.Wrap(err, "getMessage")
}

func (peer *Peer) handleMessage(msg *message.Message, currentWork []byte) ([]byte, error) {
    if msg == nil {
        // reset keep-alive
        return currentWork, nil
    }
    switch msg.ID {
    case message.MsgChoke:
        // fmt.Println("MsgChoke")
        peer.peerChoking = true
    case message.MsgUnchoke:
        // fmt.Println("MsgUnchoke")
        peer.peerChoking = false
    case message.MsgInterested:
        // fmt.Println("MsgInterested")
        peer.peerInterested = true
    case message.MsgNotInterested:
        // fmt.Println("MsgNotInterested")
        peer.peerInterested = false
    case message.MsgHave:
        // fmt.Println("MsgHave")
        if len(msg.Payload) != 4 {
            return currentWork, errors.Wrap(ErrMessage, "handleMessage")
        }
        index := binary.BigEndian.Uint32(msg.Payload)
        peer.bitfield.Set(int(index))
    case message.MsgBitfield:
        // fmt.Println("MsgBitfield")
        expected := int(math.Ceil(float64(peer.info.TotalPieces) / 8))
        if len(msg.Payload) != expected {
            return currentWork, errors.Wrap(ErrBitfield, "handleMessage")
        }
        peer.bitfield = msg.Payload
    case message.MsgRequest:
        // fmt.Println("MsgRequest")
        if len(msg.Payload) != 12 {
            return currentWork, errors.Wrap(ErrMessage, "handleMessage")
        }
        err := peer.handleRequest(msg)
        return currentWork, errors.Wrap(err, "handleMessage")
    case message.MsgPiece:
        // fmt.Println("MsgPiece")
        if len(msg.Payload) < 9 {
            return currentWork, errors.Wrap(ErrMessage, "handleMessage")
        }
        if currentWork == nil {  // Discard pieces if we are not trying to download a piece
            return nil, nil
        }
        currentWork, err := peer.handlePiece(msg, currentWork)
        return currentWork, errors.Wrap(err, "handleMessage")
    case message.MsgCancel:
        if len(msg.Payload) != 12 {
            return currentWork, errors.Wrap(ErrMessage, "handleMessage")
        }
        fmt.Println("MsgCancel not yet implemented")
    case message.MsgPort:
        if len(msg.Payload) != 2 {
            return currentWork, errors.Wrap(ErrMessage, "handleMessage")
        }
        fmt.Println("MsgPort not yet implemented")
    }
    return currentWork, nil
}

// sendRequest sends a piece request message to a peer
func (peer *Peer) sendRequest(index, begin, length int) error {
    msg := message.Request(uint32(index), uint32(begin), uint32(length))
    err := peer.Conn.Write(msg.Encode())
    return errors.Wrap(err, "sendRequest")
}

func (peer *Peer) handleRequest(msg *message.Message) error {
    if peer.amChoking {  // Tell the peer we are choking them and return
        chokeMsg := message.Choke()
        err := peer.Conn.Write(chokeMsg.Encode())
        return errors.Wrap(err, "handleRequest")
    }

    index := binary.BigEndian.Uint32(msg.Payload[0:4])
    begin := binary.BigEndian.Uint32(msg.Payload[4:8])
    length := binary.BigEndian.Uint32(msg.Payload[8:12])
    if !peer.info.Bitfield.Has(int(index)) {  // Ignore request if we don't have the piece
        return nil
    }

    piece, err := write.ReadPiece(peer.info, int(index))
    if err != nil {
        return errors.Wrap(err, "handleRequest")
    } else if len(piece) < int(begin + length) {  // Ignore request if the bounds aren't possible
        return nil
    }
    pieceMsg := message.Piece(index, begin, piece[begin:begin+length])
    err = peer.Conn.Write(pieceMsg.Encode())
    return errors.Wrap(err, "handleRequest")
}

// handlePiece adds a MsgPiece to the current work slice
func (peer *Peer) handlePiece(msg *message.Message, currentWork []byte) ([]byte, error) {
    index := binary.BigEndian.Uint32(msg.Payload[0:4])
    begin := binary.BigEndian.Uint32(msg.Payload[4:8])
    block := msg.Payload[8:]

    peer.reqsOut--
    for i := range peer.workQueue {
        if index == uint32(peer.workQueue[i].index) {
            peer.workQueue[i].left -= len(block)
        }
    }
    err := write.AddBlock(peer.info, int(index), int(begin), block, currentWork)
    return currentWork, errors.Wrap(err, "handlePiece")
}

// TODO deprecate
func (peer *Peer) getPiece(index int) ([]byte, error) {
    // Initialize peer's work
    pieceSize := common.PieceSize(peer.info, index)
    currentWork := make([]byte, pieceSize)

    // TODO start of elapsed time
    // TODO fix getting last piece (because of irregular size?)
    // for begin := 0; peer.workLeft > 0; {
    //     if !peer.peerChoking {
    //         // Send max number of requests to peer
    //         for ; peer.reqsOut < peer.rate && begin < pieceSize; {
    //             reqSize := common.Min(peer.workLeft, maxReqSize)
    //             err := peer.sendRequest(index, begin, reqSize)
    //             if err != nil {
    //                 return nil, errors.Wrap(err, "getPiece")
    //             }
    //             peer.reqsOut++
    //             begin += reqSize
    //         }
    //     }
    // }

    // TODO end of elapsed time
    return currentWork, nil
}

// downloadPiece starts a routine to download a piece from a peer
func (peer *Peer) downloadPiece(index int) ([]byte, error) {
    msg := message.Interested()
    if err := peer.Conn.Write(msg.Encode()); err != nil {
        return nil, errors.Wrap(err, "downloadPiece")
    }

    piece, err := peer.getPiece(index)
    if err != nil {
        return nil, errors.Wrap(err, "downloadPiece")
    }

    msg = message.NotInterested()
    if err := peer.Conn.Write(msg.Encode()); err != nil {
        return nil, errors.Wrap(err, "downloadPiece")
    }

    // Verify the piece's hash
    if !write.VerifyPiece(peer.info, index, piece) {
        return nil, errors.Wrap(ErrPieceHash, "downloadPiece")
    }
    return piece, nil
}

func (peer *Peer) adjustRate(actualRate uint16) {
    // Use aggressive algorithm from rtorrent
    if actualRate < 20 {
        peer.rate = actualRate + 2
    } else {
        peer.rate = actualRate / 5 + 18
    }
}
