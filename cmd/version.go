package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version     string = "1.0.0"
	betaVersion string = "1.2.0"
	beta        bool   = false
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Gives version info for Kumo",
	Long: `Kumo is a Simple CLI at its core. Just type what you wanna google and get 
	the files as HTML in your downloads folder`,

	Run: func(cmd *cobra.Command, args []string) {
		if beta {
			fmt.Println("kumo version " + betaVersion)
		} else {
			fmt.Println("kumo version " + version)
		}
	},
}

func init() {
	flags := versionCmd.Flags()
	flags.BoolVar(&beta, "beta", true, "Tells beta version")
	flags.BoolVar(&beta, "release", false, "Tells release version")
}
