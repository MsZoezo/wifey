/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
	"wifey/backends"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// messageCmd represents the message command
var messageCmd = &cobra.Command{
	Use:   "message [message]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("No message to be send, silly ;3")
		}

		message := strings.Join(args, " ")

		name := viper.GetString("general.name")

		token := viper.GetString("cloudflare.token")
		accountId := viper.GetString("cloudflare.accountId")
		queueId := viper.GetString("cloudflare.sendQueueId")

		client := backends.New(token, queueId, accountId)

		err := client.Send(backends.Message{
			Name:    name,
			Message: message,
		})

		return err
	},
}

func init() {
	rootCmd.AddCommand(messageCmd)
}
