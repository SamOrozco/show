package links

import (
	"bufio"
	"os"
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
		Name: name,
		Url:  url,
	}
}
