package cmd

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"

	pb "github.com/kylec725/gray/rpc"
	"github.com/kylec725/gray/torrent"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const pidFile = "/tmp/gray.pid"

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.AddCommand(serverMainCmd)
	serverCmd.AddCommand(serverStartCmd)
	serverCmd.AddCommand(serverStopCmd)
}

var (
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "controls the gray torrent server",
	}

	serverMainCmd = &cobra.Command{
		Use:    "main",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			session, err := torrent.NewSession()
			if err != nil {
				log.WithField("error", err).Info("Error when starting a new session for server")
			}
			// TODO: start peerListener

			// Initialize signal catching
			signalChan := make(chan os.Signal, 1)
			signal.Notify(signalChan, os.Interrupt, os.Kill, syscall.SIGTERM)
			go func() {
				_ = <-signalChan
				signal.Stop(signalChan)

				// this is a good place to flush everything to disk
				// before terminating.
				session.Close()

				// remove PID file
				os.Remove(pidFile)

				os.Exit(0)

			}()

			// Setup grpc server
			// TODO: Want to use TLS for encrypting communication
			serverAddr := ":" + strconv.Itoa(int(viper.GetInt("server.port")))
			serverListener, err := net.Listen("tcp", serverAddr)
			if err != nil {
				log.WithFields(log.Fields{"error": err.Error(), "port": serverAddr[1:]}).Fatal("Failed to listen for rpc")
			}
			server := grpc.NewServer()
			pb.RegisterTorrentServer(server, &session)
			if err = server.Serve(serverListener); err != nil {
				log.WithField("error", err).Debug("Error with serving rpc client")
			}
		},
	}

	serverStartCmd = &cobra.Command{
		Use:   "start",
		Short: "starts the gray torrent server",
		Run: func(cmd *cobra.Command, args []string) {
			// check if daemon already running.
			if _, err := os.Stat(pidFile); err == nil {
				fmt.Println("gray is already running")
				os.Exit(1)
			}

			daemon := exec.Command(os.Args[0], serverCmd.Use, serverMainCmd.Use)
			daemon.Start()
			savePID(daemon.Process.Pid)

			// fmt.Println("Daemon process ID is : ", daemon.Process.Pid)
			fmt.Println("gray started")
			os.Exit(0)

		},
	}

	serverStopCmd = &cobra.Command{
		Use:   "stop",
		Short: "stops the gray torrent server",
		Run: func(cmd *cobra.Command, args []string) {
			if _, err := os.Stat(pidFile); os.IsNotExist(err) { // do nothing if pid file does not exist
				fmt.Println("gray is not running")
				os.Exit(1)
			}

			data, err := ioutil.ReadFile(pidFile)
			if err != nil {
				fmt.Println("gray is not running")
				os.Exit(1)
			}

			pid, err := strconv.Atoi(string(data))
			if err != nil {
				fmt.Println("Unable to read and parse process ID found in ", pidFile)
				os.Exit(1)
			}

			process, err := os.FindProcess(pid)
			if err != nil {
				fmt.Printf("Unable to find process ID [%v] with error %v \n", pid, err)
				os.Exit(1)
			}
			// remove PID file
			os.Remove(pidFile)

			// fmt.Printf("Killing process ID [%v] now.\n", pid)
			// kill process and exit immediately
			err = process.Kill()
			if err != nil {
				fmt.Printf("Unable to kill process ID [%v] with error %v \n", pid, err)
				os.Exit(1)
			}

			// fmt.Printf("Killed process ID [%v]\n", pid)
			fmt.Println("gray stopped")
			os.Exit(0)
		},
	}
)

func savePID(pid int) {
	file, err := os.Create(pidFile)
	if err != nil {
		log.Fatalf("Unable to create pid file : %v\n", err)
	}
	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(pid))
	if err != nil {
		log.Fatalf("Unable to create pid file : %v\n", err)
	}
	file.Sync() // flush to disk
}