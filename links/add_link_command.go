package links

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"show_commands/utils"
)

type AddLinkCommand struct {
	linkService LinkService
}

func NewAddLinkCommand(linkService LinkService) *AddLinkCommand {
	return &AddLinkCommand{linkService: linkService}
}

func (l *AddLinkCommand) Command() *cobra.Command {
	linkCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a link",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			// get url
			fmt.Print("URL: ")
			url := utils.ReadLineFromStdIn(reader)

			// get name
			fmt.Print("Name(optional): ")
			name := utils.ReadLineFromStdIn(reader)

			link := Link{
				Url:  url,
				Name: name,
			}
			if err := l.linkService.AddLink(link); err != nil {
				fmt.Println("Error adding link: ", err)
			} else {
				fmt.Println("Link added successfully")
			}

		},
	}
	return linkCmd
}
