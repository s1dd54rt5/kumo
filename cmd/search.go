package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
	googlesearch "github.com/rocketlaunchr/google-search"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "The command to search the web and the mysteries beyond",
	Long:  `This command will search the web and put the HTML files in a path of your choice.`,
	Run: func(cmd *cobra.Command, args []string) {
		opts := &googlesearch.SearchOptions{
			Limit: 1,
		}
		results, err := googlesearch.Search(cmd.Context(), strings.Join(args[:], " "), *opts)
		if err != nil {
			fmt.Println("Error searching google")
		}
		for _, s := range results {
			resp, err := soup.Get(s.URL)
			if err != nil {
				fmt.Println("Cant fetch the site")
			}
			s1 := s.URL
			if last := len(s1) - 1; last >= 0 && s1[last] == '/' {
				s1 = s1[:last]
			}
			split := strings.Split(s1, "/")
			s1 = split[len(split)-1]
			file, err := os.Create(config.KumoPath + "/" + s1 + ".html")
			if err != nil {
				fmt.Println("Cant create file at path")
			}
			_, error := file.WriteString(resp)
			if error != nil {
				fmt.Println(err)
			}
			fmt.Println("Done!")
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
