package main

import (
	"github.com/spf13/cobra"
	"show_commands/groups"
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
	linksService := groups.NewFileSystemGroupService[*links.Link](rootDir + "/group_links_file.json")
	linkCommand := links.NewLinkCommand(linksService, links.NewLocalLinkService())
	rootCmd.AddCommand(linkCommand.Command())

	// shared flags
	rootCmd.PersistentFlags().String("group-id", "0", "This defines the group you are wanting to do your operation on. (e.g. --group-id=1 will target root group 1 --group-id=1.2 will target rootgroup 1 subgroup 2)")

	// start command
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
