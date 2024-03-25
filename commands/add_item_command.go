package commands

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"show_commands/groups"
)

type AddItemCommand[T groups.IdDisplay] struct {
	groupsService groups.GroupService[T]
	creator       func() T
}

func NewAddItemCommand[T groups.IdDisplay](groupsService groups.GroupService[T], creator func() T) *AddItemCommand[T] {
	return &AddItemCommand[T]{groupsService: groupsService, creator: creator}
}

func (l *AddItemCommand[T]) Command() *cobra.Command {
	// this command expects a group flag
	// if not provided, it will use the default group
	linkCmd := &cobra.Command{
		Use:   "add-item",
		Short: "Add a link",
		Run: func(cmd *cobra.Command, args []string) {
			// parse group id from command flag
			groupId := groups.GroupIdFromCommandString(cmd.Flags().GetString("group-id"))

			// use creator to create a new item
			item := l.creator()

			// add item to groups
			if err := l.groupsService.AddItemToGroup(groupId, item); err != nil {
				panic(err)
			}

			color.Yellow("Item added to group [%d]", groupId)
		},
	}
	return linkCmd
}
