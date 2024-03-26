package commands

import (
	"github.com/spf13/cobra"
	"show_commands/groups"
)

type ItemActionCommand[T groups.IdDisplay] struct {
	groupService groups.GroupService[T]
	actionUse    string
	actionAlias  []string
	action       func(display groups.IdDisplay) error
}

func NewItemActionCommand[T groups.IdDisplay](groupService groups.GroupService[T], actionUse string, actionAlias []string, action func(display groups.IdDisplay) error) *ItemActionCommand[T] {
	return &ItemActionCommand[T]{groupService: groupService, actionUse: actionUse, actionAlias: actionAlias, action: action}
}

func (i *ItemActionCommand[T]) Command() *cobra.Command {
	itemActionCmd := &cobra.Command{
		Use:     i.actionUse,
		Aliases: i.actionAlias,
		Run: func(cmd *cobra.Command, args []string) {
			// the arg passed to this command is going to the be the path to the link
			// e.g. 0.0.1 -> groupId 0, subgroup 0, link 1
			if len(args) == 0 {
				panic("No link provided")
				return
			}

			// get the group id
			groupId, linkId := groups.GroupIdAndItemIdFromString(args[0])
			group, err := i.groupService.GetGroupById(groupId)
			if err != nil {
				panic(err)
			}
			if linkId >= len(group.Items) {
				panic("Invalid link id")
				return
			}
			if err := i.action(group.Items[linkId]); err != nil {
				panic(err)
			}

		},
	}
	return itemActionCmd
}
