package links

import (
	"github.com/spf13/cobra"
	"show_commands/groups"
)

type OpenLinkCommand struct {
	groupService groups.GroupService[*Link]
	linkService  LinkService
}

func NewOpenLinkCommand(groupService groups.GroupService[*Link], linkService LinkService) *OpenLinkCommand {
	return &OpenLinkCommand{groupService: groupService, linkService: linkService}
}

func (l *OpenLinkCommand) Command() *cobra.Command {
	linkCmd := &cobra.Command{
		Use:     "open",
		Aliases: []string{"o"},
		Short:   "Open a link",
		Run: func(cmd *cobra.Command, args []string) {
			// the arg passed to this command is going to the be the path to the link
			// e.g. 0.0.1 -> groupId 0, subgroup 0, link 1
			if len(args) == 0 {
				panic("No link provided")
				return
			}

			// get the group id
			groupId, linkId := groups.GroupIdAndItemIdFromString(args[0])
			group, err := l.groupService.GetGroupById(groupId)
			if err != nil {
				panic(err)
			}
			if linkId >= len(group.Items) {
				panic("Invalid link id")
				return
			}
			link := group.Items[linkId]
			if err := l.linkService.OpenLink(link); err != nil {
				panic(err)
			}
		},
	}
	return linkCmd
}
