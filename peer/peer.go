/*
Package peer provides the ability to setup connections with peers as well
as manage sending and receiving torrent pieces with those peers.
Peers also handle writing pieces to file if necessary.
*/
package peer

import (
    "net"
    "strconv"
    "time"
    "math"

    "github.com/kylec725/graytorrent/common"
    "github.com/kylec725/graytorrent/bitfield"
    "github.com/kylec725/graytorrent/peer/message"
    "github.com/kylec725/graytorrent/connect"
    "github.com/pkg/errors"
    log "github.com/sirupsen/logrus"
)

const peerTimeout = 120 * time.Second
const startRate = 2  // slow approach: hard limit on requests per peer
const maxPeerQueue = 5  // Max number of pieces a peer can queue

// Peer stores info about connecting to peers as well as their state
type Peer struct {
    Host net.IP
    Port uint16
    Conn *connect.Conn  // nil if not connected

    info *common.TorrentInfo
    bitfield bitfield.Bitfield
    amChoking bool
    amInterested bool
    peerChoking bool
    peerInterested bool
    rate int  // max number of outgoing requests
    workQueue []workPiece
    shutdown bool
}

func (peer Peer) String() string {
    return net.JoinHostPort(peer.Host.String(), strconv.Itoa(int(peer.Port)))
}

// New returns a new instantiated peer
func New(host net.IP, port uint16, conn *connect.Conn, info *common.TorrentInfo) Peer {
    bitfieldSize := int(math.Ceil(float64(info.TotalPieces) / 8))
    return Peer{
        Host: host,
        Port: port,
        Conn: conn,

        info: info,
        bitfield: make([]byte, bitfieldSize),
        amChoking: true,
        amInterested: false,
        peerChoking: true,
        peerInterested: false,
        rate: startRate,
        workQueue: []workPiece{},
        shutdown: false,
    }
}

// Shutdown lets the main goroutine signal a peer to stop working
func (peer *Peer) Shutdown() {
    peer.shutdown = true
}

// Choke notifies a peer that we are choking them
func (peer *Peer) Choke() error {  // Main should handle shutting down the peer if we have an error
    peer.amChoking = true
    msg := message.Choke()
    err := peer.Conn.Write(msg.Encode())
    return errors.Wrap(err, "Choke")
}

// Unchoke notifies a peer that we are not choking them
func (peer *Peer) Unchoke() error {
    peer.amChoking = false
    msg := message.Unchoke()
    err := peer.Conn.Write(msg.Encode())
    return errors.Wrap(err, "Unchoke")
}

// StartWork makes a peer wait for pieces to download
func (peer *Peer) StartWork(work chan int, results, done chan bool) {
    peer.shutdown = false
    err := peer.verifyHandshake()
    if err != nil {
        log.WithFields(log.Fields{"peer": peer.String(), "error": err.Error()}).Debug("Handshake failed")
        // remove <- peer.String()  // Notify main to remove this peer from its list
        return
    }
    log.WithFields(log.Fields{"peer": peer.String()}).Debug("Handshake successful")

    // Setup peer connection
    connection := make(chan []byte)
    go peer.Conn.Await(connection)
    peer.Conn.Timeout = peerTimeout

    // Work loop
    for {
        // Check if main told peer to shutdown
        if peer.shutdown {
            goto exit
        }

        // TODO what happens if no data is received, but we need to get more work? i.e. this is the only peer with
        // the needed piece, but we block because we don't receive data
        select {
        case data, ok := <-connection:
            if !ok {
                goto exit
            }
            msg := message.Decode(data)
            if err = peer.handleMessage(msg, work, results); err != nil {
                // Shutdown even if error is timeout
                log.WithFields(log.Fields{"peer": peer.String(), "error": err.Error()}).Debug("Received bad message")
                goto exit
                // remove <- peer.String()  // Notify main to remove this peer from its list
            }
        case _, ok := <-done:
            if !ok {
                goto exit
            }
        default:  // TODO check if cpu usage is okay if we loop quickly with default
        }

        // Only try to find new work piece if queue is open
        if len(peer.workQueue) < peer.rate {
            select {
                // Grab work from the channel
            case index := <-work:
                // Send the work back if the peer does not have the piece
                if !peer.bitfield.Has(index) {
                    work <- index
                    continue
                }

                // Download piece from the peer
                err := peer.downloadPiece(index)
                if err != nil {
                    log.WithFields(log.Fields{"peer": peer.String(), "piece index": index, "error": err.Error()}).Debug("Starting piece download failed")
                    work <- index  // Put piece back onto work channel
                    continue
                }
            default:  // Don't block if we can't find work
            }
        }
    }

    exit:
    for i := range peer.workQueue {
        work <- peer.workQueue[i].index
    }
    peer.Conn.Quit()  // Tell connection goroutine to exit
    log.WithFields(log.Fields{"peer": peer.String()}).Debug("Peer shutdown")
    return
}
