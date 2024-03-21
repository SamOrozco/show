package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func ReadLineFromStdIn(reader *bufio.Reader) string {
	line, _, _ := reader.ReadLine()
	return string(line)
}

func CreateDirectoryIfNotExists(location string) error {
	if _, err := os.Stat(location); os.IsNotExist(err) {
		return os.MkdirAll(location, 0755)
	}
	return nil
}

func CreateFileIfNotExists(location string) error {
	if _, err := os.Stat(location); os.IsNotExist(err) {
		file, err := os.Create(location)
		if err != nil {
			return err
		}
		return file.Close()
	}
	return nil
}

func GetIntArgs(arg []string) ([]int, int) {
	if len(arg) == 0 {
		return []int{}, 0
	}

	var intArgs []int
	for _, a := range arg {
		i, err := strconv.Atoi(a)
		if err != nil {
			println("Invalid argument")
			continue
		}
		intArgs = append(intArgs, i)
	}
	return intArgs, len(intArgs)
}

func OpenLink(url string) {
	openLink(url)
}

func openLink(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
