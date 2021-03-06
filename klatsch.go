package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var cmdFetch = &cobra.Command{
		Use:   "fetch",
		Short: "fetch",
		Long:  "long fetch",
		Run:   fetch,
	}
	var Force bool
	cmdFetch.PersistentFlags().BoolVarP(&Force, "force", "F", false,
		"build HTML page even if nothing has changed")

	var cmdImportTweets = &cobra.Command{
		Use:   "import_tweets",
		Short: "import",
		Long:  "long import",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("import_tweets not implemented yet.")
		},
	}

	var cmdInit = &cobra.Command{
		Use:   "init",
		Short: "init",
		Long:  "long init",
		Run:   initStuff,
	}

	var cmdSearch = &cobra.Command{
		Use:   "search",
		Short: "search",
		Long:  "long search",
		Run:   search,
	}

	var cmdServer = &cobra.Command{
		Use:   "server",
		Short: "server",
		Long:  "long server",
		Run:   server,
	}

	var rootCmd = &cobra.Command{Use: "klatsch"}
	rootCmd.AddCommand(cmdFetch, cmdImportTweets, cmdInit, cmdSearch, cmdServer)
	rootCmd.Execute()
}
