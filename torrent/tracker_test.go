package torrent_test

import (
	"testing"

	"github.com/kylec725/graytorrent/torrent"
)

func TestTrackerPrint(t *testing.T) {
	var ta torrent.Tracker = 2
	ta.Print()
}
