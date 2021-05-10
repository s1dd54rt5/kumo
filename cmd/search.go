package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/anaskhan96/soup"
	"github.com/bharath-srinivas/termloader"
	googlesearch "github.com/rocketlaunchr/google-search"
	"github.com/spf13/cobra"
)

var listSearch = false
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "The command to search the web and the mysteries beyond",
	Long:  `This command will search the web and put the HTML files in a path of your choice.`,
	Run: func(cmd *cobra.Command, args []string) {
		if listSearch {
			opts := &googlesearch.SearchOptions{
				Limit: 3,
			}
			results, err := googlesearch.Search(cmd.Context(), strings.Join(args[:], " "), *opts)
			if err != nil {
				fmt.Println("Error searching google")
				log.Println(err)
			}
			for _, s := range results {
				fmt.Println(White + strconv.Itoa(s.Rank) + Reset + " " + Green + s.Title + Reset + "\n" + Red + s.Description + Reset + "\n")
			}
			fmt.Println("Enter the rank of article you want to download: ")
			var option int
			fmt.Scanln(&option)
			link := results[option-1].URL
			var wg sync.WaitGroup
			wg.Add(1)
			getListSite(link, &wg)
			wg.Wait()
			fmt.Println("DONE!")
		} else {
			opts := &googlesearch.SearchOptions{
				Limit: config.Search,
			}
			var wg sync.WaitGroup
			results, err := googlesearch.Search(cmd.Context(), strings.Join(args[:], " "), *opts)
			if err != nil {
				fmt.Println("Error searching google")
			}
			for _, s := range results {
				wg.Add(1)
				getSite(s, &wg)
			}
			wg.Wait()
			loader := termloader.New(termloader.CharsetConfigs["default"])
			loader.Text = "Loading ..."
			loader.Start()
			time.Sleep(time.Duration(config.Search) * 4 * time.Second)
			loader.Stop()
			fmt.Println("DONE!")
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	flags := searchCmd.Flags()
	flags.BoolVar(&listSearch, "list", true, "Tells list form of search")
}

func getSite(s googlesearch.Result, wg *sync.WaitGroup) {
	defer wg.Done()
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
}

func getListSite(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := soup.Get(url)
	if err != nil {
		fmt.Println("Cant fetch the site")
	}
	s1 := url
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
}
