package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/engine-api/client"
)

func main() {
	cli, err := client.NewEnvClient()
	quitOnErr(err)
	ver, err := cli.ServerVersion(context.Background())
	quitOnErr(err)
	fmt.Println("Was able to connect to docker daemon")
	fmt.Println(ver)
}

func quitOnErr(err error) {
	if err != nil {
		fmt.Printf("error ocurred: %v", err)
		os.Exit(1)
	}
}
