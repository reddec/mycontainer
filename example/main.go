package main

import (
	"fmt"

	"github.com/reddec/mycontainer"
)

func main() {
	containerID, err := mycontainer.ContainerID()
	if err != nil {
		panic(err)
	}

	fmt.Println("Container ID:", containerID)
}
