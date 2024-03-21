package links

import (
	"github.com/spf13/cobra"
	"strconv"
)

type SwapLinkCommand struct {
	linkService LinkService
}

func NewSwapLinkCommand(linkService LinkService) *SwapLinkCommand {
	return &SwapLinkCommand{linkService: linkService}
}

func (l *SwapLinkCommand) Command() *cobra.Command {
	linkCmd := &cobra.Command{
		Use:   "swap",
		Short: "Swap a link",
		Long:  "Switch the indexes of two links",
		Run: func(cmd *cobra.Command, args []string) {
			// get id
			id1 := 0
			id2 := 0
			var err error
			if len(args) > 0 {
				id1, err = strconv.Atoi(args[0])
				if err != nil {
					panic(err)
				}

				id2, err = strconv.Atoi(args[1])
				if err != nil {
					panic(err)
				}
			}

			if err := l.linkService.SwapLink(id1, id2); err != nil {
				panic(err)
			}

			println("Link swapped successfully")
		},
	}
	return linkCmd
}
