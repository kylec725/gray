package main

import (
	"os"
    "fmt"

    "github.com/kylec725/graytorrent/connect"
    "github.com/kylec725/graytorrent/torrent"
    flag "github.com/spf13/pflag"
    log "github.com/sirupsen/logrus"
    viper "github.com/spf13/viper"
)

const logLevel = log.TraceLevel  // InfoLevel || DebugLevel || TraceLevel

var (
    logFile *os.File
    port uint16
    err error
    filename string
    torrentList []torrent.Torrent
)

func init() {
    flag.StringVarP(&filename, "file", "f", "", "Filename of torrent file")
    flag.Parse()

    setupLog()
    log.Info("Graytorrent started")

    setupViper()
    viper.WatchConfig()
}

func init() {
    portRange := viper.GetIntSlice("network.portrange")
    port, err = connect.OpenPort(portRange)
    if err != nil {
        log.WithFields(log.Fields{
            "portrange": portRange,
            "error": err.Error(),
        }).Warn("No open port found in portrange, using random port")
        // TODO get a random port to use for the client
    }
}

// Initialize GUI
func init() {

}

func main() {
    defer logFile.Close()
    // defer g.Close()

    // Handle single torrent download for now
    if filename != "" {
        addTorrent(filename)
    }

    // Send torrent stopped messages
    // Save torrent progresses to history file

    log.Info("Graytorrent stopped")
}

func addTorrent(filename string) {
    to := torrent.Torrent{Path: filename}
    if err := to.Setup(); err != nil {
        log.WithFields(log.Fields{"file": filename, "error": err.Error()}).Info("Torrent setup failed")
        return
    }
    log.WithField("name", to.Info.Name).Info("Torrent added")
    torrentList = append(torrentList, to)
    to.Start()
    // to.Shutdown()  // signal for a torrent to shutdown
}
