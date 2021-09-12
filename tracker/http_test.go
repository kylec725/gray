package tracker

import (
	"fmt"
	"testing"

	"github.com/kylec725/graytorrent/common"
	"github.com/kylec725/graytorrent/metainfo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const debugHTTP = false

func TestBuildURL(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	meta, err := metainfo.Meta("../tmp/1056.txt.utf-8.torrent")
	require.Nil(err, "Error with metainfo.Meta()")
	info, err := common.GetInfo(meta)
	require.Nil(err, "GetInfo() error")

	trackers, err := GetTrackers(meta)

	for _, tr := range trackers {
		assert.NotNil(tr)
		// Build url with each tracker object
		url, err := tr.buildURL("started", info, uint16(6881), 0, 0, info.Left)
		if err == nil {
			if debugTracker {
				fmt.Println("HTTP tracker request:", url)
			}
			assert.Equal("http://poole.cs.umd.edu:6969/announce?compact=1&downloaded=0&event=started&info_hash=Q%CB%DD%21%F2FYx%DAc%F0%91%B1y%18g2%CCX%05&left=810121&numwant=30&peer_id="+string(info.PeerID[:])+"&port=6881&uploaded=0", url, "Built tracker request was incorrect")
		}
	}
}

func TestHTTPReqs(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	meta, err := metainfo.Meta("../tmp/batonroad.torrent")
	require.Nil(err, "Error with metainfo.Meta()")
	info, err := common.GetInfo(meta)
	require.Nil(err, "GetInfo() error")

	trackers, err := GetTrackers(meta)

	var testTracker Tracker
	for _, tr := range trackers {
		assert.NotNil(tr)

		if tr.Announce[0:4] == "http" {
			testTracker = tr
		}
	}

	if debugHTTP {
		fmt.Printf("HTTP Tracker%+v\n", testTracker)
	}

	peersList, err := testTracker.sendStarted(info, 6881, 0, 0, info.Left)
	if err != nil { // Resend because this tracker fails on first retry for no reason
		peersList, err = testTracker.sendStarted(info, 6881, 0, 0, info.Left)
	}
	if assert.Nil(err) {
		for _, peer := range peersList {
			if debugHTTP {
				fmt.Println("HTTP Peer:", peer)
			}
		}
		err = testTracker.sendStopped(info, 6881, 0, 0, info.Left)
		assert.Nil(err)

		if debugHTTP {
			fmt.Printf("HTTP Tracker%+v\n", testTracker)
		}
	}
}
