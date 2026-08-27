package cmd

import (
	"wifey/internal/socket"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(socketCmd)
}

var socketCmd = &cobra.Command{
	Use:   "socket",
	Short: "Start wifey's socket process",
	RunE: func(cmd *cobra.Command, args []string) error {
		server, err := socket.InitServer()

		if err != nil {
			return err
		}

		server.InitSignalHandler()

		server.Handle()

		return nil
	},
}
