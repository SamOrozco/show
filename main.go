package main

import (
	"github.com/spf13/cobra"
	"os"
	command_line "show_commands/command-line"
	"show_commands/commands"
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

var rootDir = getRootDir()

func main() {
	if err := utils.CreateDirectoryIfNotExists(rootDir); err != nil {
		panic(err)
	}
	// link command
	rootCmd.AddCommand(buildLinksCommand())
	// command line command
	rootCmd.AddCommand(buildCommandLineCommand())

	// shared flags
	rootCmd.PersistentFlags().StringP("group-id", "g", "0", "This defines the group you are wanting to do your operation on. (e.g. --group-id=1 will target root group 1 --group-id=1.2 will target rootgroup 1 subgroup 2)")

	// start command
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func buildCommandLineCommand() *cobra.Command {
	// command line service
	commandLineService := groups.NewFileSystemGroupService[*command_line.CommandLine](rootDir + "/group_command_line_file.json")
	commandLineCommand := command_line.NewCommandLineCommand(commandLineService)
	commandLineCmd := commandLineCommand.Command()
	commandLineCmd.AddCommand(commands.NewItemActionCommand(commandLineService, "copy", []string{"c"}, func(display groups.IdDisplay) error {
		commandLine := display.(*command_line.CommandLine)
		return utils.CopyToClipboard(commandLine.Code)
	}).Command())
	return commandLineCmd
}

func buildLinksCommand() *cobra.Command {
	linksGroupService := groups.NewFileSystemGroupService[*links.Link](rootDir + "/group_links_file.json")
	linkService := links.NewLocalLinkService()

	linkCommand := links.NewLinkCommand(linksGroupService)
	linkCmd := linkCommand.Command()
	linkCmd.AddCommand(commands.NewItemActionCommand(linksGroupService, "open", []string{"o"},
		func(display groups.IdDisplay) error {
			link := display.(*links.Link)
			return linkService.OpenLink(link)
		}).Command())
	return linkCmd
}

func getRootDir() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return userHomeDir + "/.show"
}
