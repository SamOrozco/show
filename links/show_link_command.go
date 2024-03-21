package links

import "github.com/spf13/cobra"

type ShowLinkCommand struct {
	linkService LinkService
}

func NewShowLinkCommand(linkService LinkService) *ShowLinkCommand {
	return &ShowLinkCommand{linkService: linkService}
}

func (l *ShowLinkCommand) Command() *cobra.Command {
	linkCmd := &cobra.Command{
		Use:   "show",
		Short: "Show a link",
		Run: func(cmd *cobra.Command, args []string) {
			links, err := l.linkService.GetLinks()
			if err != nil {
				panic(err)
			}
			l.linkService.PrintLinks(links)
		},
	}
	return linkCmd
}
