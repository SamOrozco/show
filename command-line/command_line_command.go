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

func (c *CommandLine) GetName() string {
	return c.Name
}

func (c *CommandLine) GetId() int {
	return c.Id
}

func (c *CommandLine) SetId(id int) {
	c.Id = id
}

var titleFormatter = color.New(color.FgHiGreen).SprintfFunc()
var codeFormatter = color.New().SprintfFunc()

func (c *CommandLine) DisplayString(parentId string) string {
	var strBldr strings.Builder
	strBldr.WriteString(titleFormatter(fmt.Sprintf("[%s.%d][%s]", parentId, c.Id, c.Name)))
	strBldr.WriteString(titleFormatter(fmt.Sprintf(" -> ")))
	strBldr.WriteString(utils.TruncateString(codeFormatter(fmt.Sprintf(" %s ", c.Code)), 50, 200))
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
			if len(args) > 0 {
				currentGroups = filterGroupBySearchText(currentGroups, args[0])
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

// filterGroupBySearchText filters the group list by the search text
func filterGroupBySearchText(groupList []*groups.Group[*CommandLine], searchText string) []*groups.Group[*CommandLine] {
	if searchText == "" {
		return groupList
	}
	returnValue := make([]*groups.Group[*CommandLine], 0)
	for _, group := range groupList {
		if strings.Contains(group.Name, searchText) {
			returnValue = append(returnValue, group)
		} else if items := filterItemsBySearchText(group, searchText); len(items) > 0 {
			group.Items = items
			returnValue = append(returnValue, group)
		}
	}
	return returnValue
}

// filterItemsBySearchText filters the items in a group by the search text
func filterItemsBySearchText(group *groups.Group[*CommandLine], searchText string) []*CommandLine {
	if searchText == "" {
		return group.Items
	}
	resultItems := make([]*CommandLine, 0)
	for _, item := range group.Items {
		if strings.Contains(item.GetName(), searchText) {
			resultItems = append(resultItems, item)
		}
	}
	return resultItems
}
