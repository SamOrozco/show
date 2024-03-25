package commands

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"show_commands/groups"
	"show_commands/utils"
)

type AddSubGroupCommand[T groups.IdDisplay] struct {
	groupService groups.GroupService[T]
}

func NewAddSubGroupCommand[T groups.IdDisplay](groupService groups.GroupService[T]) *AddSubGroupCommand[T] {
	return &AddSubGroupCommand[T]{groupService: groupService}
}

func (l *AddSubGroupCommand[T]) Command() *cobra.Command {
	linkCmd := &cobra.Command{
		Use:     "add-sub-group",
		Aliases: []string{"asg"},
		Short:   "Add a sub group",
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

			// get group from command
			// if no group provided will be added to default group
			groupId := groups.GroupIdFromCommandString(cmd.Flags().GetString("group-id"))

			// add sub command
			if err := l.groupService.AddSubGroup(groupId, &groups.Group[T]{
				Name:      groupName,
				Items:     []T{},
				SubGroups: []*groups.Group[T]{},
			}); err != nil {
				panic(err)
			}
			color.Yellow("Sub-Group added!")
		},
	}
	return linkCmd
}
