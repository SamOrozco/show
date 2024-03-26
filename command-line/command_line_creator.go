package command_line

import (
	"bufio"
	"os"
	"show_commands/utils"
)

func NewCommandLineCreator() *CommandLine {
	lineReader := bufio.NewReader(os.Stdin)

	print("Name: ")
	name, err := lineReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	print("Code: ")
	code, err := lineReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return &CommandLine{
		Name: utils.CleanString(name),
		Code: utils.CleanString(code),
	}
}
