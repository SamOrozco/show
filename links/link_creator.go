package links

import (
	"bufio"
	"os"
	"show_commands/utils"
)

func NewLinkCreator() *Link {
	lineReader := bufio.NewReader(os.Stdin)

	print("URL: ")
	url, err := lineReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	print("Name: ")
	name, err := lineReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return &Link{
		Name: utils.CleanString(name),
		Url:  utils.CleanString(url),
	}
}
