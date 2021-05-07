package cmd

import (
	"fmt"
	"strings"

	googlesearch "github.com/rocketlaunchr/google-search"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		results, err := googlesearch.Search(cmd.Context(), strings.Join(args[:], " "))
		if err != nil {
			fmt.Println("Error searching google")
		}
		for _, s := range results {
			fmt.Println(s.URL)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
