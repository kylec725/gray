package main

import (
	"io"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// ErrListener is used to close the peerListener safely
var ErrListener = errors.New("use of closed network connection")

func initLog() {
	// Logging file
	logFile, err = os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// logFile, err = os.OpenFile(filepath.Join(grayTorrentPath, "info.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.WithField("error", err.Error()).Fatal("Could not open log file")
	}

	// Set logging settings
	log.SetOutput(logFile)
	log.SetFormatter(&prefixed.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})
	if verbose {
		dualOutput := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(dualOutput)
	}
	log.SetLevel(logLevel)
}

func initConfig() {
	viper.SetDefault("torrent.path", ".")
	viper.SetDefault("torrent.autoseed", true)
	viper.SetDefault("network.portrange", [2]int{6881, 6889})
	viper.SetDefault("network.connections.globalMax", 300)
	viper.SetDefault("network.connections.torrentMax", 30)

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	// viper.AddConfigPath(".") // Remove in the future
	viper.AddConfigPath(grayTorrentPath)
	viper.AddConfigPath("/etc/graytorrent")

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, create default config
			viper.SafeWriteConfig()
			log.Info("Config file written")
		} else {
			// Some other error was found
			log.Panic("Fatal error reading config file:", err)
		}
	}
}

// func catchInterrupt(ctx context.Context, cancel context.CancelFunc) {
// 	signalChan := make(chan os.Signal, 1)
// 	signal.Notify(signalChan, os.Interrupt)
// 	select {
// 	case <-signalChan: // Cleanup on interrupt signal
// 		signal.Stop(signalChan)
// 		peerListener.Close()
// 		cancel()
// 		err = torrent.SaveAll(torrentList)
// 		if err != nil {
// 			log.WithField("error", err).Debug("Problem occurred while saving torrent management data")
// 		}
// 		log.Info("Graytorrent stopped")
// 		logFile.Close()
// 		os.Exit(1)
// 	case <-ctx.Done():
// 	}
// }
