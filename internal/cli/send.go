package cmd

import (
	"fmt"
	"strings"
	"wifey/internal/socket"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sendCmd)
}

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message to your wifeys",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := socket.InitClient()

		if err != nil {
			return err
		}

		defer client.Close()

		str := fmt.Sprintf(`send "%s"`, strings.Join(args, " "))

		res, err := client.Send(str)

		if err != nil {
			return err
		}

		fmt.Println(res)

		return nil
	},
}
