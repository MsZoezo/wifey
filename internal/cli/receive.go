package cmd

import (
	"fmt"
	"wifey/internal/socket"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(receiveCmd)
}

var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Maybe there is a message for you on the wire!",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := socket.InitClient()

		if err != nil {
			return err
		}

		defer client.Close()

		res, err := client.Send("receive")

		if err != nil {
			return err
		}

		fmt.Println(res)

		return nil
	},
}
