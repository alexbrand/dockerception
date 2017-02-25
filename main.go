package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Echo this from within the a container
const dc = `
______               _                                    _    _               
|  _  \             | |                                  | |  (_)              
| | | |  ___    ___ | | __  ___  _ __   ___   ___  _ __  | |_  _   ___   _ __  
| | | | / _ \  / __|| |/ / / _ \| '__| / __| / _ \| '_ \ | __|| | / _ \ | '_ \ 
| |/ / | (_) || (__ |   < |  __/| |   | (__ |  __/| |_) || |_ | || (_) || | | |
|___/   \___/  \___||_|\_\ \___||_|    \___| \___|| .__/  \__||_| \___/ |_| |_|
                                                  | |                          
						                                                                             
`

func main() {
	fmt.Printf("\nMax procs = %d\n\n", runtime.GOMAXPROCS(-1))
	runOrQuit("docker", "version")
	runOrQuit("docker", "pull", "busybox")
	runOrQuit("docker", "run", "busybox", "echo", dc)
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
