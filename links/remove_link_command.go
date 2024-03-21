package links

import (
	"github.com/spf13/cobra"
	"strconv"
)

type RemoveLinkCommand struct {
	linkService LinkService
}

func NewRemoveLinkCommand(linkService LinkService) *RemoveLinkCommand {
	return &RemoveLinkCommand{linkService: linkService}
}

func (l *RemoveLinkCommand) Command() *cobra.Command {
	linkCmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a link",
		Run: func(cmd *cobra.Command, args []string) {
			// get id
			id := 0
			var err error
			if len(args) > 0 {
				id, err = strconv.Atoi(args[0])
				if err != nil {
					panic(err)
				}
			}

			if err := l.linkService.RemoveLink(id); err != nil {
				panic(err)
			}
		},
	}
	return linkCmd
}
