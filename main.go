package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	cmd := exec.Command("rclone",
		"--progress",
		"copy",
		// local 2gb file to copy
		"/Users/lakhansamani/Movies/Archive.zip",
		"dest/copy_backup.zip")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

	err = cmd.Start()
	fmt.Println("The command is running")
	if err != nil {
		fmt.Println(err)
	}

	// print the output of the subprocess
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		outputText := scanner.Text()
		if strings.HasPrefix(outputText, " *") {
			regex, err := regexp.Compile("\\d+\\%")
			if err != nil {
				panic(err)
			}
			matches := regex.FindAllString(outputText, -1)
			fmt.Println("=> Percentage Progress:", matches[1])
		}
	}
	cmd.Wait()
}
