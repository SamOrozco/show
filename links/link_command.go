package links

import (
	"github.com/spf13/cobra"
	"show_commands/commands"
	"show_commands/groups"
)

type LinkCommand struct {
	groupService groups.GroupService[*Link]
}

func NewLinkCommand(groupService groups.GroupService[*Link]) *LinkCommand {
	return &LinkCommand{groupService: groupService}
}

func (l *LinkCommand) Command() *cobra.Command {
	linkCmd := &cobra.Command{
		Use:     "links",
		Aliases: []string{"l"},
		Short:   "Manage links, show and add links to manage",
		Run: func(cmd *cobra.Command, args []string) {
			currentGroups, err := l.groupService.GetGroups()
			if err != nil {
				panic(err)
			}
			l.groupService.PrintGroups(currentGroups)
		},
	}
	linkCmd.AddCommand(commands.NewShowCommand(l.groupService).Command())
	linkCmd.AddCommand(commands.NewAddItemCommand(l.groupService, NewLinkCreator).Command())
	linkCmd.AddCommand(commands.NewAddGroupCommand(l.groupService).Command())
	linkCmd.AddCommand(commands.NewAddSubGroupCommand(l.groupService).Command())

	//linkCmd.AddItemCommand(NewRemoveLinkCommand(l.linkService).Command())
	//linkCmd.AddItemCommand(NewSwapLinkCommand(l.linkService).Command())
	//linkCmd.AddItemCommand(NewOpenLinkCommand(l.linkService).Command())
	return linkCmd
}
