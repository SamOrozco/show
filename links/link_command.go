package links

import "github.com/spf13/cobra"

type LinkCommand struct {
	linkService LinkService
}

func NewLinkCommand(linkService LinkService) *LinkCommand {
	return &LinkCommand{linkService: linkService}
}

func (l *LinkCommand) Command() *cobra.Command {
	linkCmd := &cobra.Command{
		Use:     "links",
		Aliases: []string{"l"},
		Short:   "Manage links, show and add links to manage",
		Run: func(cmd *cobra.Command, args []string) {
			links, err := l.linkService.GetLinks()
			if err != nil {
				panic(err)
			}
			l.linkService.PrintLinks(links)
		},
	}
	linkCmd.AddCommand(NewAddLinkCommand(l.linkService).Command())
	linkCmd.AddCommand(NewShowLinkCommand(l.linkService).Command())
	linkCmd.AddCommand(NewRemoveLinkCommand(l.linkService).Command())
	linkCmd.AddCommand(NewSwapLinkCommand(l.linkService).Command())
	linkCmd.AddCommand(NewOpenLinkCommand(l.linkService).Command())
	return linkCmd
}
