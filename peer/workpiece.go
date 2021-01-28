package peer

import (
	"time"

	"github.com/kylec725/graytorrent/common"
)

type workPiece struct {
	index     int
	piece     []byte
	left      int // bytes remaining in piece
	curr      int // current byte position in slice
	startTime time.Time
}

func (peer *Peer) addWorkPiece(info common.TorrentInfo, index int) {
	pieceSize := common.PieceSize(info, index)
	piece := make([]byte, pieceSize)
	newWork := workPiece{index, piece, pieceSize, 0, time.Now()}
	peer.workQueue = append(peer.workQueue, newWork)
}

func (peer *Peer) removeWorkPiece(index int) {
	removeIndex := -1
	for i, workPiece := range peer.workQueue {
		if index == workPiece.index {
			removeIndex = i
		}
	}
	if removeIndex != -1 {
		peer.workQueue[removeIndex] = peer.workQueue[len(peer.workQueue)-1]
		peer.workQueue = peer.workQueue[:len(peer.workQueue)-1]
	}
}

// clearWork sends peer's work back into the work pool
func (peer *Peer) clearWork(work chan int) {
	for _, wp := range peer.workQueue {
		work <- wp.index
	}
	peer.workQueue = peer.workQueue[:0]
}
