package main

import (
	"github.com/spf13/cobra"
	"show_commands/links"
	"show_commands/utils"
)

var (
	rootCmd = &cobra.Command{
		Use:   "show",
		Short: "A command line tool for managing commands and links",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

const rootDir = "/tmp/show/links"

func main() {
	if err := utils.CreateDirectoryIfNotExists(rootDir); err != nil {
		panic(err)
	}

	// link service
	linksService := links.NewLocalLinkService(rootDir + "/links_file.json")
	linkCommand := links.NewLinkCommand(linksService)
	rootCmd.AddCommand(linkCommand.Command())

	// start command
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
