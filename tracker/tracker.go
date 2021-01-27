package tracker

import (
    "time"
    "math/rand"
    "net/http"

    "github.com/kylec725/graytorrent/common"
    "github.com/kylec725/graytorrent/metainfo"
    "github.com/kylec725/graytorrent/peer"
    "github.com/pkg/errors"
    log "github.com/sirupsen/logrus"
)

// TODO figure out way to get bytes left in torrent to trackers accurately

// Errors
var (
    ErrNoAnnounce = errors.New("Did not find any annouce urls")
)

// Tracker stores information about a torrent tracker
type Tracker struct {
    Announce string
    Working bool
    Interval int

    info *common.TorrentInfo
    httpClient *http.Client
    port uint16
}

func newTracker(announce string, info *common.TorrentInfo, port uint16) Tracker {
    return Tracker{
        Announce: announce,
        Working: false,
        Interval: 60,

        info: info,
        httpClient: &http.Client{ Timeout: 20 * time.Second },
        port: port,
    }
}

// GetTrackers parses metainfo to retrieve a list of trackers
func GetTrackers(meta metainfo.BencodeMeta, info *common.TorrentInfo, port uint16) ([]Tracker, error) {
    // If announce-list is empty, use announce only
    if len(meta.AnnounceList) == 0 {
        // Check if no announce strings exist
        if meta.Announce == "" {
            return nil, errors.Wrap(ErrNoAnnounce, "getTrackers")
        }

        trackers := make([]Tracker, 1)
        trackers[0] = newTracker(meta.Announce, info, port)
        return trackers, nil
    }

    // Make list of multiple trackers
    var trackers []Tracker
    var numAnnounce int
    // Add each announce in announce-list as a tracker
    for _, group := range meta.AnnounceList {
        for _, announce := range group {
            trackers = append(trackers, newTracker(announce, info, port))
            numAnnounce++
        }
    }

    // Shuffle list before returning
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(numAnnounce, func(x, y int) {
        trackers[x], trackers[y] = trackers[y], trackers[x]
    })

    return trackers, nil
}

// Run starts a tracker and gets peers for a torrent
func (tr *Tracker) Run(peers chan peer.Peer, done chan bool) {
    ctxLog := log.WithField("tracker", tr.Announce)
    peerList, err := tr.sendStarted()  // hardcoded number of bytes left
    if err != nil {
        tr.Working = false
        ctxLog.WithField("error", err.Error()).Debug("Error while sending started message")
    } else {
        tr.Working = true
        ctxLog.WithField("amount", len(peerList)).Debug("Received list of peers")
    }

    // Cleanup
    defer func() {
        if tr.Working {  // Send stopped message if necessary
            if err = tr.sendStopped(); err != nil {
                ctxLog.WithField("error", err.Error()).Debug("Error while sending stopped message")
            }
        }
    }()

    // Send peers through channel
    for i := range peerList {
        peers <- peerList[i]
    }

    for {
        select {
        case _, ok := <-done:
            if !ok {
                return
            }
        case <-time.After(time.Duration(tr.Interval) * time.Second):
            // Contact tracker again
        // default:
        //     // TODO try to connect to tracker again after interval
        //     if !tr.Working {
        //
        //     }
        }
    }
}
