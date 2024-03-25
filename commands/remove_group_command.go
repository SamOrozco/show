package commands

import (
	"github.com/spf13/cobra"
	"show_commands/groups"
)

type RemoveGroupCommand[T groups.IdDisplay] struct {
	groupService groups.GroupService[T]
}

func NewRemoveGroupCommand[T groups.IdDisplay](groupService groups.GroupService[T]) *RemoveGroupCommand[T] {
	return &RemoveGroupCommand[T]{groupService: groupService}
}

func (r *RemoveGroupCommand[T]) Command() *cobra.Command {
	removeGroupCmd := &cobra.Command{
		Use:     "remove",
		Aliases: []string{"r"},
		Short:   "Remove a group",
		Run: func(cmd *cobra.Command, args []string) {
			err := r.groupService.RemoveGroup(groups.GroupIdFromCommandString(cmd.Flags().GetString("group-id")))
			if err != nil {
				panic(err)
			}
		},
	}
	return removeGroupCmd
}
