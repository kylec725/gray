/*
Package peer provides the ability to setup connections with peers as well
as manage sending and receiving torrent pieces with those peers.
Peers also handle writing pieces to file if necessary.
*/
package peer

import (
	"context"
	"math"
	"net"
	"time"

	"github.com/kylec725/graytorrent/bitfield"
	"github.com/kylec725/graytorrent/common"
	"github.com/kylec725/graytorrent/connect"
	"github.com/kylec725/graytorrent/peer/message"
	log "github.com/sirupsen/logrus"
)

const peerTimeout = 20 * time.Second       // Time to wait on an expected peer connection operation
const requestTimeout = 15 * time.Second    // How long to wait on requests before sending work back
const keepAliveTimeout = 120 * time.Second // How long to wait before removing a peer with no messages
const sendKeepAlive = 90 * time.Second     // How long to wait before sending a keep alive message
const adjustTime = 5 * time.Second         // How often in seconds to adjust the transfer rates

// Peer stores info about connecting to peers as well as their state
type Peer struct {
	Addr           string
	Conn           *connect.Conn // nil if not connected
	AmChoking      bool
	AmInterested   bool
	PeerChoking    bool
	PeerInterested bool
	Send           chan message.Message // Used by outer goroutines to send messages, allows us to handle errors internally

	bitfield     bitfield.Bitfield
	workPieces   map[int]workPiece // Map to keep track of what pieces we're trying to get
	queue        int               // How many requests have been sent out
	queueSize    int               // How many requests can be queued at a time
	bytesRcvd    uint32            // Number of bytes received since the last adjustment time
	bytesSent    uint32            // Number of bytes sent since the last adjustment time
	lastMsgRcvd  time.Time
	lastMsgSent  time.Time
	lastRequest  time.Time // Last time a request was sent
	lastPiece    time.Time // Last time a piece was received
	lastUnchoked time.Time
}

func (p Peer) String() string {
	return p.Addr
}

// New returns a new instantiated peer
func New(addr string, conn net.Conn, info common.TorrentInfo) Peer {
	var peerConn *connect.Conn = nil
	if conn != nil {
		peerConn = &connect.Conn{Conn: conn, Timeout: peerTimeout}
	}
	bitfieldSize := int(math.Ceil(float64(info.TotalPieces) / 8))
	return Peer{
		Addr:           addr,
		Conn:           peerConn,
		AmChoking:      true,
		AmInterested:   false,
		PeerChoking:    true,
		PeerInterested: false,
		Send:           make(chan message.Message),

		bitfield:    make([]byte, bitfieldSize),
		workPieces:  make(map[int]workPiece),
		queue:       0,
		queueSize:   minQueue,
		bytesRcvd:   0,
		bytesSent:   0,
		lastMsgRcvd: time.Now(),
		lastMsgSent: time.Now(),
		lastRequest: time.Now(),
		lastPiece:   time.Now(),
	}
}

// StartWork makes a peer wait for pieces to download
func (p *Peer) StartWork(ctx context.Context, work chan int, results chan int, deadPeers chan string) {
	info := common.Info(ctx)
	peerLog := log.WithField("peer", p.String())
	if p.Conn == nil {
		if err := p.dial(); err != nil {
			peerLog.WithField("error", err.Error()).Debug("Dial failed")
			deadPeers <- p.String() // Notify main to remove this peer from its list
			return
		} else if err := p.initHandshake(info); err != nil {
			peerLog.WithField("error", err.Error()).Debug("Handshake failed")
			deadPeers <- p.String()
			return
		}
	}
	peerLog.Debug("Handshake successful")

	// Setup peer connection
	connCtx, connCancel := context.WithCancel(ctx)
	p.Conn.Timeout = peerTimeout
	connection := make(chan []byte, 2) // Buffer so that connection can exit if we haven't read the data yet
	go p.Conn.Poll(connCtx, connection)

	// Create ticker to update the adaptive queuing rate
	adapRateTicker := time.NewTicker(adjustTime)

	// Cleanup
	defer func() {
		deadPeers <- p.String() // Notify main to remove this peer from its list
		p.clearWork(work)
		connCancel()
		adapRateTicker.Stop()
		peerLog.Debug("Peer shutdown")
	}()

	// Work loop
	for {
		select {
		case <-ctx.Done():
			return
		case data, ok := <-connection: // Incoming data from peer
			if !ok { // Connection failed
				return
			}
			p.lastMsgRcvd = time.Now()
			currInfo := common.Info(ctx)
			msg := message.Decode(data)
			if err := p.handleMessage(msg, currInfo, work, results); err != nil {
				peerLog.WithFields(log.Fields{"type": msg.String(), "size": len(msg.Payload), "error": err.Error()}).Debug("Error handling message")
				return
			}
		case msg := <-p.Send:
			if err := p.sendMessage(&msg); err != nil {
				peerLog.WithFields(log.Fields{"type": msg.String(), "error": err.Error()}).Debug("Error sending message")
				return
			}
		case <-adapRateTicker.C:
			p.adjustRate()
			if p.lastRequest.Sub(p.lastPiece) >= requestTimeout {
				p.clearWork(work)
				msg := message.NotInterested()
				if err := p.sendMessage(&msg); err != nil {
					peerLog.WithFields(log.Fields{"type": msg.String(), "error": err.Error()}).Debug("Error sending message")
					return
				}
			}
			if time.Since(p.lastMsgSent) >= sendKeepAlive {
				msg := (*message.Message)(nil)
				if err := p.sendMessage(msg); err != nil {
					peerLog.WithFields(log.Fields{"type": msg.String(), "error": err.Error()}).Debug("Error sending message")
					return
				}
			}
			if time.Since(p.lastMsgRcvd) >= keepAliveTimeout { // Check if peer has passed the keep-alive time
				return
			}
		}

		// Find new work piece if queue is open
		if p.queue < p.queueSize {
			select {
			case index := <-work:
				// Send the work back if the peer does not have the piece
				if !p.bitfield.Has(index) {
					work <- index
					continue
				}

				// Download piece from the peer
				err := p.downloadPiece(info, index)
				if err != nil {
					peerLog.WithFields(log.Fields{"piece index": index, "error": err.Error()}).Debug("Failed to start piece download")
					return
				}
			default: // Don't block if we can't find work
			}
		}
		if !p.PeerChoking {
			err := p.fillQueue()
			if err != nil {
				peerLog.WithField("error", err.Error()).Debug("Issue when filling queue")
				return
			}
		}
	}
}
