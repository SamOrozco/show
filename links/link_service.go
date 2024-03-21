package links

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/glamour"
	"os"
	"show_commands/utils"
	"strings"
)

type Link struct {
	Id   int
	Name string
	Url  string
}

func (l Link) String() string {
	if l.Name == "" {
		l.Name = "No name"
	}
	return fmt.Sprintf(`- [%d]%s %s`, l.Id, l.Name, l.Url)
}

type LinkService interface {
	AddLink(link Link) error
	GetLinks() ([]*Link, error)
	RemoveLink(id int) error
	PrintLinks(links []*Link)
	SwapLink(idx1, idx2 int) error
	GetLink(id int) (*Link, error)
	OpenLink(link *Link) error
}

type localLinkService struct {
	rootFileLocation string
}

func NewLocalLinkService(rootFileLocation string) LinkService {
	return &localLinkService{rootFileLocation: rootFileLocation}
}

func (l localLinkService) AddLink(link Link) error {
	if err := utils.CreateFileIfNotExists(l.rootFileLocation); err != nil {
		return err
	}

	links, err := l.readAllLinks()
	if err != nil {
		return err
	}
	// ADDING LINKS TO LIST
	links = append(links, &link)

	// update ids for links
	for i, link := range links {
		link.Id = i
	}

	if err := l.writeLinksToFile(links); err != nil {
		return err
	}
	return nil
}

func (l localLinkService) GetLinks() ([]*Link, error) {
	if err := utils.CreateFileIfNotExists(l.rootFileLocation); err != nil {
		return nil, err
	}
	return l.readAllLinks()
}

func (l localLinkService) PrintLinks(links []*Link) {
	if len(links) == 0 {
		println("No links found")
		return
	}

	var bldr strings.Builder
	bldr.WriteString("# Links\n")
	for i := range links {
		bldr.WriteString(links[i].String() + "\n\n")
	}

	output, err := glamour.Render(bldr.String(), "dark")
	if err != nil {
		panic(err)
	}
	fmt.Print(output)
}

func (l localLinkService) RemoveLink(id int) error {
	println("Removing link with id: ", id)

	links, err := l.readAllLinks()
	if err != nil {
		return err
	}

	if id < 0 || id >= len(links) {
		return nil
	}

	// remove link from list
	links = append(links[:id], links[id+1:]...)

	// update ids for links
	for i, link := range links {
		link.Id = i
	}

	if err := l.writeLinksToFile(links); err != nil {
		return err
	}
	return nil
}

func (l localLinkService) SwapLink(idx1, idx2 int) error {
	links, err := l.readAllLinks()
	if err != nil {
		return err
	}

	if idx1 < 0 || idx1 >= len(links) || idx2 < 0 || idx2 >= len(links) {
		println("Invalid indexes")
		return nil
	}

	links[idx1], links[idx2] = links[idx2], links[idx1]

	for i, link := range links {
		link.Id = i
	}

	if err := l.writeLinksToFile(links); err != nil {
		return err
	}
	return nil

}

func (l localLinkService) GetLink(id int) (*Link, error) {
	links, err := l.readAllLinks()
	if err != nil {
		return nil, err
	}

	if id < 0 || id >= len(links) {
		return nil, nil
	}

	return links[id], nil
}

func (l localLinkService) OpenLink(link *Link) error {
	if link == nil {
		println("Invalid link")
		return nil
	}
	utils.OpenLink(link.Url)
	return nil
}

func (l localLinkService) writeLinksToFile(links []*Link) error {
	data, err := json.Marshal(links)
	if err != nil {
		return err
	}
	if err := os.WriteFile(l.rootFileLocation, data, 0644); err != nil {
		return err
	}
	return nil
}

func (l localLinkService) readAllLinks() ([]*Link, error) {
	data, err := os.ReadFile(l.rootFileLocation)
	if err != nil {
		return nil, err
	}
	// if this is the first link we need to return an empty list
	if len(data) == 0 {
		return make([]*Link, 0), nil
	}
	links := make([]*Link, 0)
	if err := json.Unmarshal(data, &links); err != nil {
		return nil, err
	}
	return links, nil
}
