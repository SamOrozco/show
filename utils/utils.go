package utils

import (
	"bufio"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func PromptFromStdIn(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	return ReadLineFromStdIn(reader)
}

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

func CleanString(value string) string {
	return strings.Trim(strings.TrimSpace(value), "\n")
}

func PromptForConfirmation(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	fmt.Println("Enter 'y' to confirm or any other key to cancel")
	fmt.Print(">")
	line, _, _ := reader.ReadLine()
	return strings.ToLower(string(line)) == "y"
}

func CopyToClipboard(value string) error {
	color.Yellow("Copying to clipboard")
	return clipboard.WriteAll(value)
}

func ReadFromClipboard() (string, error) {
	return clipboard.ReadAll()
}

func TruncateString(value string, length int) string {
	if len(value) > length {
		return value[:length]
	}
	return value
}
