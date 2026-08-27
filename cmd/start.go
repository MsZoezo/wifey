package cmd

import (
	"wifey/core"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start and listen on unix socket",
	RunE: func(cmd *cobra.Command, args []string) error {
		socket, err := core.InitSocket()

		if err != nil {
			return err
		}

		socket.InitSignalHandler()
		socket.Handle()

		return nil
	},
}
