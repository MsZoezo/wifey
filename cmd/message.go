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

		fmt.Printf("Message: '%v'\n", message)

		token := viper.GetString("cloudflare.token")
		accountId := viper.GetString("cloudflare.accountId")
		queueId := viper.GetString("cloudflare.queueId")

		fmt.Println(token)

		client := backends.New(token, queueId, accountId)

		if client == nil {
			fmt.Println("AAAAA")
		}

		err := client.Send(message)

		return err
	},
}

func init() {
	rootCmd.AddCommand(messageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// messageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// messageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
