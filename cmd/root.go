package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wifey",
	Short: "Send & receive messages to/from your nerdy wifey (or anyone really!) :3",
	// TODO: Add a long description for wifey. :3
	// Long: ``,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return initConfig(cmd)
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/wifey/config.toml)")

	rootCmd.PersistentFlags().String("name", "", "Sets your name, be silly ;3")
	viper.BindPFlag("general.name", rootCmd.PersistentFlags().Lookup("name"))
}

func initConfig(cmd *cobra.Command) error {
	viper.SetEnvPrefix("WIFEY")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "*", "-", "*"))
	viper.AutomaticEnv()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/wifey/")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return fmt.Errorf("Config file is required.")
		} else {
			return err
		}
	}

	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return err
	}

	return nil
}
