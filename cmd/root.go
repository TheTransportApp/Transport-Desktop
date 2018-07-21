package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/spf13/viper"
	"github.com/mitchellh/go-homedir"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:	"transport",
	Short:	"(S)FTP and SSH client",
	Long:	"This application is a modern (S)FTP and SSH client based on Golang",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Launching UI ...")
		// Launch UI
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.transport-desktop.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".transport-desktop")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}
}