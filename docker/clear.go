package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/fatih/color"
)

func ClearLocal() {
	color.Cyan("Nana Init")
	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("%v %v\n", " => Connect docker engine:", color.RedString("Failed"))
		fmt.Println(" => Can not connect docker engine - " + err.Error())
		return
	}

	//clear docker image not use
	if _, err := cli.ImagesPrune(context.Background(), filters.NewArgs()); err != nil {
		fmt.Printf("%v %v\n", " => Clear Docker Images:", color.RedString("Failed"))
		fmt.Println(" => Error - " + err.Error())
		return
	}
	fmt.Printf("%v %v\n", " => Clear Docker Images:", color.GreenString("Success"))

	//clear docker container exited
	if _, err := cli.ContainersPrune(context.Background(), filters.NewArgs()); err != nil {
		fmt.Printf("%v %v\n", " => Clear Docker Containers:", color.RedString("Failed"))
		fmt.Println(" => Error - " + err.Error())
		return
	}
	fmt.Printf("%v %v\n", " => Clear Docker Containers:", color.GreenString("Success"))

	//clear docker volume not use
	if _, err := cli.VolumesPrune(context.Background(), filters.NewArgs()); err != nil {
		fmt.Printf("%v %v\n", " => Clear Docker Volumns:", color.RedString("Failed"))
		fmt.Println(" => Error - " + err.Error())
		return
	}
	fmt.Printf("%v %v\n", " => Clear Docker Volumns:", color.GreenString("Success"))

	//clear docker network not use
	if _, err := cli.NetworksPrune(context.Background(), filters.NewArgs()); err != nil {
		fmt.Printf("%v %v\n", " => Clear Docker Network:", color.RedString("Failed"))
		fmt.Println(" => Error - " + err.Error())
		return
	}
	fmt.Printf("%v %v\n", " => Clear Docker Network:", color.GreenString("Success"))
}
