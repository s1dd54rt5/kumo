package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type Config struct {
	KumoPath string
	Search   int
}

var config Config

var rootCmd = &cobra.Command{
	Use:   "kumo",
	Short: "Kumo helps you get files from a line",
	Long: `Kumo is a Simple CLI at its core. Just type what you wanna google and get 
	the files as HTML in your downloads folder`,

	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := homedir.Dir()
	cobra.CheckErr(err)
	viper.AddConfigPath(home)
	viper.SetConfigName(".kumo")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error in locating config file")
	}
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Error in unmarshalling")
	}
}
