package links

import (
	"bufio"
	"os"
	"show_commands/utils"
)

func NewLinkCreator(readLinkFromClipboard bool) *Link {
	lineReader := bufio.NewReader(os.Stdin)
	print("Name: ")
	name, err := lineReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	var url string
	if readLinkFromClipboard {
		url, err = utils.ReadFromClipboard()
	} else {
		print("URL: ")
		url, err = lineReader.ReadString('\n')
	}
	if err != nil {
		panic(err)
	}
	return &Link{
		Name: utils.CleanString(name),
		Url:  utils.CleanString(url),
	}
}
