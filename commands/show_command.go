package commands

import (
	"github.com/spf13/cobra"
	"show_commands/groups"
)

type ShowCommand[T groups.IdDisplay] struct {
	groupsService groups.GroupService[T]
}

func NewShowCommand[T groups.IdDisplay](groupService groups.GroupService[T]) *ShowCommand[T] {
	return &ShowCommand[T]{groupsService: groupService}
}

func (l *ShowCommand[T]) Command() *cobra.Command {
	linkCmd := &cobra.Command{
		Use:   "show",
		Short: "Show a link",
		Run: func(cmd *cobra.Command, args []string) {
			links, err := l.groupsService.GetGroups()
			if err != nil {
				panic(err)
			}
			l.groupsService.PrintGroups(links)
		},
	}
	return linkCmd
}
