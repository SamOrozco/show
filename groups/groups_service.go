package groups

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"os"
	"show_commands/utils"
	"strconv"
	"strings"
)

type GroupId struct {
	Data []int
}

// GroupIdFromCommandString will parse the flag value and return a GroupId struct
func GroupIdFromCommandString(flagValue string, err error) *GroupId {
	// if there an error getting the flag value return default group
	if err != nil {
		return &GroupId{Data: []int{0}}
	}
	return GroupIdFromString(flagValue)
}

func GroupIdFromString(flagValue string) *GroupId {
	// if no value provided we are going to default to the default group
	if flagValue == "" {
		return &GroupId{Data: []int{0}}
	}

	// we know there are child groups to parse out
	if strings.Contains(flagValue, ".") {
		var data []int
		for _, s := range strings.Split(flagValue, ".") {
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			data = append(data, i)
		}
		return &GroupId{Data: data}
	}

	// else single root group
	i, err := strconv.Atoi(flagValue)
	if err != nil {
		panic(err)
	}
	return &GroupId{Data: []int{i}}
}

func GroupIdAndItemIdFromString(value string) (*GroupId, int) {
	groupId := GroupIdFromString(value)
	return groupId, groupId.PopLast()
}

func (g *GroupId) HasSubGroups() bool {
	return len(g.Data) > 1
}

func (g *GroupId) GetRootId() int {
	return g.Data[0]
}

func (g *GroupId) GetSubGroup(idx int) int {
	return g.Data[idx]
}

func (g *GroupId) String() string {
	var strBuilder strings.Builder
	for i, d := range g.Data {
		strBuilder.WriteString(strconv.Itoa(d))
		if i != len(g.Data)-1 {
			strBuilder.WriteString(".")
		}
	}
	return strBuilder.String()
}

// PopLast will remove the last element from the group id
// and return the last value that was popped
func (g *GroupId) PopLast() int {
	last := g.Data[len(g.Data)-1]
	g.Data = g.Data[:len(g.Data)-1]
	return last
}

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

// output func
type FormatFunc func(format string, a ...interface{}) string

var headerDepthFunc = []FormatFunc{
	color.New(color.FgHiBlue).SprintfFunc(),
	color.New(color.FgHiYellow).SprintfFunc(),
	color.New(color.FgHiGreen).SprintfFunc(),
	color.New(color.FgHiMagenta).SprintfFunc(),
	color.New(color.FgHiCyan).SprintfFunc(),
	color.New(color.FgHiRed).SprintfFunc(),
	color.New(color.FgHiWhite).SprintfFunc(),
	color.New(color.FgHiBlue).SprintfFunc(),
}

var indexString = " "

func (g *Group[T]) DisplayString(depth int) string {
	var headerFormatFunc FormatFunc
	if depth > len(headerDepthFunc)-1 {
		headerFormatFunc = headerDepthFunc[0]
	} else {
		headerFormatFunc = headerDepthFunc[depth]
	}

	// GROUP HEADER
	var strBuilder strings.Builder
	strBuilder.WriteString(strings.Repeat(indexString, depth) + headerFormatFunc("[%d]%s", g.Id, g.Name) + "\n")

	// WRITE ITEMS
	for _, item := range g.Items {
		strBuilder.WriteString(strings.Repeat(indexString, depth+2) + item.DisplayString() + "\n")
	}

	// SUB GROUPS
	if len(g.SubGroups) > 0 {
		for _, subGroup := range g.SubGroups {
			strBuilder.WriteString(strings.Repeat(indexString, depth+1) + subGroup.DisplayString(depth+1) + "\n")
		}
	}
	return strBuilder.String()
}

type GroupService[T IdDisplay] interface {
	GetGroups() ([]*Group[T], error)
	AddGroup(group *Group[T]) error
	AddSubGroup(groupId *GroupId, group *Group[T]) error
	GetGroupByName(name string) (*Group[T], error)
	GetGroupById(id *GroupId) (*Group[T], error)
	AddItemToGroup(groupId *GroupId, item T) error
	RemoveItemFromGroup(groupId *GroupId, itemId int) error
	RemoveGroup(id *GroupId) error
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
		bldr.WriteString(group.DisplayString(0) + "\n")
	}
	fmt.Print(bldr.String())
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

func (f fileSystemGroupService[T]) AddSubGroup(groupId *GroupId, subGroup *Group[T]) error {
	groups, err := f.GetGroups()
	if err != nil {
		return err
	}
	// get root subGroup
	group := groups[groupId.GetRootId()]
	if groupId.HasSubGroups() {
		for i := 1; i < len(groupId.Data); i++ {
			group = group.SubGroups[groupId.GetSubGroup(i)]
		}
	}
	// add subGroup to group
	group.SubGroups = append(group.SubGroups, subGroup)

	// set ids
	for i, g := range group.SubGroups {
		g.Id = i
	}

	data, err := json.Marshal(groups)
	if err != nil {
		return err
	}
	return os.WriteFile(f.filePath, data, 0644)

}

func (f fileSystemGroupService[T]) RemoveGroup(id *GroupId) error {
	groups, err := f.GetGroups()
	if err != nil {
		return err
	}

	// get root group
	hasSubGroups := id.HasSubGroups()
	if !hasSubGroups {
		newGroups := []*Group[T]{}
		for i, group := range groups {
			if i != id.GetRootId() {
				newGroups = append(newGroups, group)
			}
		}
		for i, group := range newGroups {
			group.Id = i
		}
		data, err := json.Marshal(newGroups)
		if err != nil {
			return err
		}
		return os.WriteFile(f.filePath, data, 0644)
	} else {
		println("Removing sub groups is not supported yet")
		return nil
	}
}

func (f *fileSystemGroupService[T]) AddItemToGroup(groupId *GroupId, item T) error {
	groups, err := f.GetGroups()
	if err != nil {
		return err
	}

	// get root group
	group := groups[groupId.GetRootId()]
	if groupId.HasSubGroups() {
		for i := 1; i < len(groupId.Data); i++ {
			group = group.SubGroups[groupId.GetSubGroup(i)]
		}
	}

	newItems := append(group.Items, item)
	for i, newItem := range newItems {
		newItem.SetId(i)
	}
	group.Items = newItems
	data, err := json.Marshal(groups)
	if err != nil {
		return err
	}
	return os.WriteFile(f.filePath, data, 0644)
}

func (f *fileSystemGroupService[T]) RemoveItemFromGroup(groupId *GroupId, itemId int) error {
	groups, err := f.GetGroups()
	if err != nil {
		return err
	}
	group := groups[groupId.GetRootId()]
	if groupId.HasSubGroups() {
		for i := 1; i < len(groupId.Data); i++ {
			group = group.SubGroups[groupId.GetSubGroup(i)]
		}
	}

	var newItems []T
	for i, item := range group.Items {
		if i != itemId {
			newItems = append(newItems, item)
		}
	}

	for i, newItem := range newItems {
		newItem.SetId(i)
	}
	group.Items = newItems
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

func (f *fileSystemGroupService[T]) GetGroupById(id *GroupId) (*Group[T], error) {
	groups, err := f.GetGroups()
	if err != nil {
		return nil, err
	}
	// get root group
	group := groups[id.GetRootId()]
	if id.HasSubGroups() {
		for i := 1; i < len(id.Data); i++ {
			group = group.SubGroups[id.GetSubGroup(i)]
		}
	}
	return group, nil
}
