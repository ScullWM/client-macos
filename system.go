package main

import (
	"fmt"
	"log"
	"os/exec"
)

func openDirectory(p string) {
	cmd := exec.Command("open", p)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func updateScreencaptureDirectory(p string) {
	destinationCmd := fmt.Sprint("defaults write com.apple.screencapture location " + p)
	cmd := exec.Command("sh", "-c", destinationCmd)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	cmdKill := exec.Command("killall SystemUIServer")
	err2 := cmdKill.Start()
	if err != nil {
		log.Fatal(err2)
	}
}
