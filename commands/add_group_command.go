package commands

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"show_commands/groups"
	"show_commands/utils"
)

type AddGroupCommand[T groups.IdDisplay] struct {
	groupService groups.GroupService[T]
}

func NewAddGroupCommand[T groups.IdDisplay](groupService groups.GroupService[T]) *AddGroupCommand[T] {
	return &AddGroupCommand[T]{groupService: groupService}
}

func (l *AddGroupCommand[T]) Command() *cobra.Command {
	linkCmd := &cobra.Command{
		Use:     "add-group",
		Aliases: []string{"ag"},
		Short:   "Add a group",
		Run: func(cmd *cobra.Command, args []string) {
			groupName := ""
			if len(args) < 1 {
				groupName = utils.PromptFromStdIn("Enter group name: ")
			} else {
				groupName = args[0]
			}
			if groupName == "" {
				println("Invalid group name")
				return
			}
			if err := l.groupService.AddGroup(&groups.Group[T]{
				Name:      groupName,
				Items:     []T{},
				SubGroups: []*groups.Group[T]{},
			}); err != nil {
				panic(err)
			}

			color.Yellow("Group added!")
		},
	}
	return linkCmd
}
