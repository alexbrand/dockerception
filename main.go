package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("docker", "run", "busybox", "echo", "Dockerception!")
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
