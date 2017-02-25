package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"github.com/docker/engine-api/types/strslice"
)

func main() {
	// Create docker client
	cli, err := client.NewEnvClient()
	quitOnErr(err)

	// Let's see if we can even talk to Docker daemon
	ver, err := cli.ServerVersion(context.Background())
	quitOnErr(err)
	fmt.Println("Was able to connect to docker daemon")
	fmt.Println(ver)

	// Let's create a container
	cmd := strslice.StrSlice{"echo", "hello", "world"}
	containerConf := &container.Config{
		Image: "busybox",
		Cmd:   cmd,
	}
	resp, err := cli.ContainerCreate(context.Background(), containerConf, nil, nil, "")
	quitOnErr(err)

	// Let's start the container
	err = cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
	quitOnErr(err)
	fmt.Println("Was able to run a container")
}

func quitOnErr(err error) {
	if err != nil {
		fmt.Printf("error ocurred: %v", err)
		os.Exit(1)
	}
}
