package links

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"show_commands/utils"
	"strconv"
	"strings"
)

type OpenLinkCommand struct {
	linkService LinkService
}

func NewOpenLinkCommand(linkService LinkService) *OpenLinkCommand {
	return &OpenLinkCommand{linkService: linkService}
}

func (l *OpenLinkCommand) Command() *cobra.Command {
	linkCmd := &cobra.Command{
		Use:   "open",
		Short: "Open a link",
		Run: func(cmd *cobra.Command, args []string) {
			intArgs, size := utils.GetIntArgs(args)
			// if an arg is passed we are going to automatically open the link with that Id
			// if no arg is passed we are going to prompt the user to open a link

			linkId := 0
			showDialog := false
			if size < 1 {
				showDialog = true
			} else {
				linkId = intArgs[0]
			}
			if err := l.OpenLink(linkId, showDialog); err != nil {
				panic(err)
			}

		},
	}
	return linkCmd
}

func (l *OpenLinkCommand) OpenLink(openId int, showDialog bool) error {
	if !showDialog {
		link, err := l.linkService.GetLink(openId)
		if err != nil {
			return err
		}
		if err := l.linkService.OpenLink(link); err != nil {
			return err
		}
		return nil
	}

	links, err := l.linkService.GetLinks()
	if err != nil {
		return err
	}
	l.linkService.PrintLinks(links)
	fmt.Print("Open a link: ")
	line := utils.ReadLineFromStdIn(bufio.NewReader(os.Stdin))
	linkId, err := strconv.Atoi(strings.TrimSpace(line))
	if err := l.linkService.OpenLink(links[linkId]); err != nil {
		return err
	}
	return nil
}
