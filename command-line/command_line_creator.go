package command_line

import (
	"bufio"
	"os"
	"show_commands/utils"
)

func NewCommandLineCreator(readInputFromClipboard bool) *CommandLine {
	lineReader := bufio.NewReader(os.Stdin)
	print("Name: ")
	name, err := lineReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	var code string
	if readInputFromClipboard {
		code, err = utils.ReadFromClipboard()
		print("Read value from clipboard")
	} else {
		print("Code: ")
		code, err = lineReader.ReadString('\n')
	}
	if err != nil {
		panic(err)
	}
	return &CommandLine{
		Name: utils.CleanString(name),
		Code: utils.CleanString(code),
	}
}
