package cmd

import (
	"fmt"

	"github.com/kylec725/graytorrent/client"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&magnetLink, "magnet", "m", "", "use a magnet link instead of a .torrent file to add a torrent")
	rootCmd.AddCommand(listCmd)
}

var (
	listCmd = &cobra.Command{
		Use:   "ls",
		Short: "list the currently managed torrents",
		Run: func(cmd *cobra.Command, args []string) {
			if err := client.List(); err != nil {
				fmt.Println(err)
			}
		},
	}

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "adds a new torrent from a .torrent file or magnet link",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := client.Add(args[0]); err != nil {
				fmt.Println(err)
			}
		},
	}

	removeCmd = &cobra.Command{
		Use:   "rm",
		Short: "removes a managed torrent",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := client.Remove(args[0]); err != nil {
				fmt.Println(err)
			}
		},
	}

	startCmd = &cobra.Command{
		Use:   "start",
		Short: "starts a torrent's download/upload",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := client.Start(args[0]); err != nil {
				fmt.Println(err)
			}
		},
	}

	stopCmd = &cobra.Command{
		Use:   "start",
		Short: "starts a torrent's download/upload",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := client.Stop(args[0]); err != nil {
				fmt.Println(err)
			}
		},
	}
)
