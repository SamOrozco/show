package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"show_commands/groups"
	"show_commands/utils"
)

type RemoveItemCommand[T groups.IdDisplay] struct {
	groupService groups.GroupService[T]
}

func NewRemoveItemCommand[T groups.IdDisplay](groupService groups.GroupService[T]) *RemoveItemCommand[T] {
	return &RemoveItemCommand[T]{groupService: groupService}
}

func (r *RemoveItemCommand[T]) Command() *cobra.Command {
	removeItemCmd := &cobra.Command{
		Use:     "remove-item",
		Aliases: []string{"ri"},
		Short:   "Remove an item",
		Run: func(cmd *cobra.Command, args []string) {
			// the arg passed to this command is going to the be the path to the item
			// e.g. 0.0.1 -> groupId 0, subgroup 0, item 1
			if len(args) == 0 {
				panic("No item provided")
				return
			}

			// get the group id
			groupId, itemId := groups.GroupIdAndItemIdFromString(args[0])
			if !utils.PromptForConfirmation(fmt.Sprintf("Are you sure you want to remove item [%d] from group [%s]?", itemId, groupId)) {
				return
			}

			if err := r.groupService.RemoveItemFromGroup(groupId, itemId); err != nil {
				panic(err)
			}
			fmt.Println("Item removed")
		},
	}
	return removeItemCmd
}
