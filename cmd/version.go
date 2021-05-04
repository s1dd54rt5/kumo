package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version string = "1.0.0"
)
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Gives version info for Kumo",
	Long: `Kumo is a Simple CLI at its core. Just type what you wanna google and get 
	the files as HTML in your downloads folder`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kumo version " + version)
	},
}
