package docker

import (
	"context"
	"log"

	"github.com/docker/docker/api/types/network"
)

func CreateNetwork(name string) string {
	resp, err := cli.NetworkCreate(context.Background(), name, network.CreateOptions{})
	if err != nil {
		log.Default().Println(err)
	}

	return resp.ID
}

func RemoveNetwork(name string) {
	err := cli.NetworkRemove(context.Background(), name)
	if err != nil {
		log.Default().Println(err)
	}
}
