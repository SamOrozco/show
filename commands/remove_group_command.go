package commands

import (
	"fmt"
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
		Use:     "remove-group",
		Aliases: []string{"rg"},
		Short:   "Remove a group",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				panic("No group provided")
				return
			}
			groupId := groups.GroupIdFromString(args[0])
			err := r.groupService.RemoveGroup(groupId)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Group %d removed", groupId)
		},
	}
	return removeGroupCmd
}
