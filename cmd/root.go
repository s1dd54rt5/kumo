package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kumo",
	Short: "Kumo is a CLI build in go to fetch HTML Pages.",
	Long:  "Kumo is a CLI build in go to fetch HTML Pages. It can be useful to quickly get data from the net.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Success!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
