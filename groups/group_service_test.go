package groups

import (
	"show_commands/links"
	"testing"
)

func TestFileSystemGroupService_AddGroup(t *testing.T) {
	group := Group[links.Link]{
		Name:      "Base group",
		Items:     []links.Link{{Id: 1, Name: "Link 1", Url: "http://link1.com"}},
		SubGroups: []Group[links.Link]{{Name: "Sub group", Items: []links.Link{{Id: 2, Name: "Link 2", Url: "http://link2.com"}}}},
	}

	linkGroupService := NewFileSystemGroupService[links.Link]("/tmp/groups.json")
	if err := linkGroupService.AddGroup(&group); err != nil {
		t.Fatal(err)
	}
}

func TestFileSystemGroupService_GetGroups(t *testing.T) {
	linkGroupService := NewFileSystemGroupService[links.Link]("/tmp/groups.json")
	groups, err := linkGroupService.GetGroups()
	if err != nil {
		t.Fatal(err)
	}
	if len(groups) == 0 {
		t.Fatal("No groups found")
	}
	for i := range groups {
		t.Logf("Group: %+v\n", groups[i])
	}
}
