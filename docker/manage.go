package docker

import (
	"context"

	"github.com/docker/docker/api/types/container"
)

func RunImage(image string) error {
	err := PullImage(image)
	if err != nil {
		return err
	}

	resp, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: image,
	}, nil, nil, nil, "")
	if err != nil {
		return err
	}

	return cli.ContainerStart(context.Background(), resp.ID, container.StartOptions{})
}

func StopContainer(id string) error {
	return cli.ContainerStop(context.Background(), id, container.StopOptions{})
}
