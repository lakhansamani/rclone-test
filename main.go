package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// Prepare rclone command
	c := exec.Command("rclone",
		"--progress",
		"copy",
		"src/test_file.txt",
		"dest/test_file.txt",
	)

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err := c.Run()
	if err != nil {
		log.Fatal("=> err:", err)
	}
}
