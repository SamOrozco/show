package links

import (
	"fmt"
	"show_commands/utils"
)

type Link struct {
	Id   int
	Name string
	Url  string
}

func (l *Link) GetName() string {
	return l.Name
}

func (l *Link) GetId() int {
	return l.Id
}

func (l *Link) SetId(id int) {
	l.Id = id
}

func (l *Link) DisplayString(parentID string) string {
	return l.String(parentID)
}

func (l *Link) String(parentId string) string {
	if l.Name == "" {
		l.Name = "No name"
	}
	return fmt.Sprintf(`- [%s.%d]%s %s`, parentId, l.Id, l.Name, l.Url)
}

type LinkService interface {
	OpenLink(link *Link) error
}

type localLinkService struct {
}

func NewLocalLinkService() LinkService {
	return &localLinkService{}
}

func (l localLinkService) OpenLink(link *Link) error {
	if link == nil {
		println("Invalid link")
		return nil
	}
	utils.OpenLink(link.Url)
	return nil
}
