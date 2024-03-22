package groups

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/glamour"
	"os"
	"show_commands/utils"
	"strings"
)

type IdDisplay interface {
	GetId() int
	SetId(id int)
	DisplayString() string
}

type Group[T IdDisplay] struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	Items     []T         `json:"items"`
	SubGroups []*Group[T] `json:"sub_groups"`
}

func (g *Group[T]) GetId() int {
	return g.Id
}

func (g *Group[T]) SetId(id int) {
	g.Id = id
}

func (g *Group[T]) DisplayString() string {

	// GROUP HEADER
	var strBuilder strings.Builder
	strBuilder.WriteString("# " + fmt.Sprintf("[%d]%s", g.Id, g.Name) + "\n")

	// WRITE ITEMS
	for _, item := range g.Items {
		strBuilder.WriteString(item.DisplayString() + "\n")
	}

	// SUB GROUPS
	strBuilder.WriteString(fmt.Sprintf("# %s -> SubGroups", g.Name) + "\n")
	for _, subGroup := range g.SubGroups {
		strBuilder.WriteString(subGroup.DisplayString() + "\n")
	}
	return strBuilder.String()
}

type GroupService[T IdDisplay] interface {
	GetGroups() ([]*Group[T], error)
	AddGroup(group *Group[T]) error
	AddSubGroup(groupId int, group *Group[T]) error
	GetGroupByName(name string) (*Group[T], error)
	GetGroupById(id int) (*Group[T], error)
	AddItemToGroup(groupId int, item T) error
	RemoveItemFromGroup(groupId int, itemId int) error
	RemoveGroup(id int) error
	PrintGroups(groups []*Group[T])
}

type fileSystemGroupService[T IdDisplay] struct {
	filePath string
}

func NewFileSystemGroupService[T IdDisplay](filePath string) GroupService[T] {
	if err := utils.CreateFileIfNotExists(filePath); err != nil {
		panic(err)
	}
	return &fileSystemGroupService[T]{filePath: filePath}
}

func (f *fileSystemGroupService[T]) GetGroups() ([]*Group[T], error) {
	data, err := os.ReadFile(f.filePath)
	if err != nil {
		return nil, err
	}
	// if this is the first time adding a group we are going to return a list with a default group
	if len(data) == 0 {
		return []*Group[T]{{
			Name:      "Default Group",
			Items:     []T{},
			SubGroups: []*Group[T]{},
		}}, nil
	}
	var groups []*Group[T]
	if err = json.Unmarshal(data, &groups); err != nil {
		return nil, err
	}
	return groups, nil
}

func (f *fileSystemGroupService[T]) PrintGroups(groups []*Group[T]) {
	if len(groups) < 1 {
		fmt.Println("No groups found")
		return
	}

	var bldr strings.Builder
	for _, group := range groups {
		bldr.WriteString(group.DisplayString() + "\n")
	}
	out, err := glamour.Render(bldr.String(), "dark")
	if err != nil {
		panic(err)
	}
	fmt.Print(out)
}

func (f fileSystemGroupService[T]) AddGroup(group *Group[T]) error {
	groups, err := f.GetGroups()
	if err != nil {
		return err
	}
	// add to group and reset ids
	groups = append(groups, group)
	for i, g := range groups {
		g.Id = i
	}
	data, err := json.Marshal(groups)
	if err != nil {
		return err
	}
	return os.WriteFile(f.filePath, data, 0644)
}

func (f fileSystemGroupService[T]) AddSubGroup(groupId int, group *Group[T]) error {
	groups, err := f.GetGroups()
	if err != nil {
		return err
	}
	groups[groupId].SubGroups = append(groups[groupId].SubGroups, group)

	// set ids
	for i, g := range groups[groupId].SubGroups {
		g.Id = i
	}

	data, err := json.Marshal(groups)
	if err != nil {
		return err
	}
	return os.WriteFile(f.filePath, data, 0644)

}

func (f fileSystemGroupService[T]) RemoveGroup(id int) error {
	groups, err := f.GetGroups()
	if err != nil {
		return err
	}
	// remove group and reset ids
	var newGroups []*Group[T]
	for i, g := range groups {
		if i != id {
			newGroups = append(newGroups, g)
		}
	}
	for i, g := range newGroups {
		g.Id = i
	}
	data, err := json.Marshal(newGroups)
	if err != nil {
		return err
	}
	return os.WriteFile(f.filePath, data, 0644)
}

func (f *fileSystemGroupService[T]) AddItemToGroup(groupId int, item T) error {
	groups, err := f.GetGroups()
	if err != nil {
		return err
	}
	newItems := append(groups[groupId].Items, item)
	for i, newItem := range newItems {
		newItem.SetId(i)
	}
	groups[groupId].Items = newItems
	data, err := json.Marshal(groups)
	if err != nil {
		return err
	}
	return os.WriteFile(f.filePath, data, 0644)
}

func (f *fileSystemGroupService[T]) RemoveItemFromGroup(groupId int, itemId int) error {
	groups, err := f.GetGroups()
	if err != nil {
		return err
	}
	// remove item from group
	var newItems []T
	for i, item := range groups[groupId].Items {
		if i != itemId {
			newItems = append(newItems, item)
		}
	}

	// set id on new items
	for i, item := range newItems {
		item.SetId(i)
	}

	groups[groupId].Items = newItems
	data, err := json.Marshal(groups)
	if err != nil {
		return err
	}
	return os.WriteFile(f.filePath, data, 0644)
}

func (f *fileSystemGroupService[T]) GetGroupByName(name string) (*Group[T], error) {
	groups, err := f.GetGroups()
	if err != nil {
		return nil, err
	}
	for _, group := range groups {
		if group.Name == name {
			return group, nil
		}
	}
	return nil, nil
}

func (f *fileSystemGroupService[T]) GetGroupById(id int) (*Group[T], error) {
	groups, err := f.GetGroups()
	if err != nil {
		return nil, err
	}
	if id < 0 || id >= len(groups) {
		return nil, nil
	}
	return groups[id], nil
}
