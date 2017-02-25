package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	runOrQuit("docker", "pull", "busybox")
	runOrQuit("docker", "run", "busybox", "echo", "Dockerception!")
	runOrQuit("docker", "ps", "-a")
}

func runOrQuit(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	quitOnErr(err)
}

func quitOnErr(err error) {
	if err != nil {
		fmt.Printf("error ocurred: %v", err)
		os.Exit(1)
	}
}
