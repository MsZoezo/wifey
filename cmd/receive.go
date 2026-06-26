/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"wifey/backends"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// receiveCmd represents the receive command
var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		token := viper.GetString("cloudflare.token")
		accountId := viper.GetString("cloudflare.accountId")
		queueId := viper.GetString("cloudflare.receiveQueueId")

		client := backends.New(token, queueId, accountId)

		message, err := client.Receive()

		if err != nil {
			return err
		}

		if message != nil {
			fmt.Println(message.Message)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(receiveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// receiveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// receiveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
