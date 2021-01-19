package main

import (
	"os"

    "github.com/kylec725/graytorrent/connect"
    flag "github.com/spf13/pflag"
    log "github.com/sirupsen/logrus"
    viper "github.com/spf13/viper"
)

const debug = false

var (
    logFile *os.File
    port uint16
    err error
    filename string
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
        log.WithField("portrange", portRange).Warn("No open port found in portrange, using random port")
        // TODO get a random port to use for the client
    }
}

// Initialize GUI
func init() {

}

func main() {
    defer logFile.Close()
    // defer g.Close()

    // Send torrent stopped messages
    // Save torrent progresses to history file

    log.Info("Graytorrent ended")
}

