package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

const version = "0.0.3+git"

func printVersions(binarydir string) {
	fmt.Println("etcd-starter version", version)

	binary1 := path.Join(binarydir, "1", "etcd")
	binary2 := path.Join(binarydir, "2", "etcd")

	if _, err := os.Stat(binary1); err == nil {
		output, err := exec.Command(binary1, "--version").Output()
		if err == nil {
			fmt.Print("internal version 1: " + string(output))
		}
	}

	if _, err := os.Stat(binary2); err == nil {
		output, err := exec.Command(binary2, "--version").Output()
		if err == nil {
			fmt.Print("internal version 2: " + string(output))
		}
	}
}
