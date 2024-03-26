package command_line

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"show_commands/commands"
	"show_commands/groups"
	"show_commands/utils"
	"strings"
)

type CommandLine struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func (c *CommandLine) GetId() int {
	return c.Id
}

func (c *CommandLine) SetId(id int) {
	c.Id = id
}

var titleFormatter = color.New(color.FgHiGreen).SprintfFunc()
var codeFormatter = color.New(color.BgWhite).SprintfFunc()

func (c *CommandLine) DisplayString() string {
	var strBldr strings.Builder
	strBldr.WriteString(titleFormatter(fmt.Sprintf("[%d][%s]", c.Id, c.Name)))
	strBldr.WriteString(titleFormatter(fmt.Sprintf(" -> ")))
	strBldr.WriteString(utils.TruncateString(codeFormatter(fmt.Sprintf(" %s ", c.Code)), 100))
	return strBldr.String()
}

type CommandLineCommand struct {
	groupsService groups.GroupService[*CommandLine]
}

func NewCommandLineCommand(groupsService groups.GroupService[*CommandLine]) *CommandLineCommand {
	return &CommandLineCommand{groupsService: groupsService}
}

func (c *CommandLineCommand) Command() *cobra.Command {
	command := &cobra.Command{
		Use:     "command-line",
		Short:   "Manage command line snippets",
		Aliases: []string{"cli", "c"},
		Run: func(cmd *cobra.Command, args []string) {
			currentGroups, err := c.groupsService.GetGroups()
			if err != nil {
				panic(err)
			}
			c.groupsService.PrintGroups(currentGroups)
		},
	}
	command.AddCommand(commands.NewShowCommand(c.groupsService).Command())
	command.AddCommand(commands.NewAddItemCommand(c.groupsService, NewCommandLineCreator).Command())
	command.AddCommand(commands.NewAddGroupCommand(c.groupsService).Command())
	command.AddCommand(commands.NewAddSubGroupCommand(c.groupsService).Command())
	command.AddCommand(commands.NewRemoveGroupCommand(c.groupsService).Command())
	command.AddCommand(commands.NewRemoveItemCommand(c.groupsService).Command())

	return command
}
