package docker

import (
	"context"

	"github.com/docker/docker/api/types/container"
)

func RunImage(image string) (*string, error) {
	err := PullImage(image)
	if err != nil {
		return nil, err
	}

	resp, err := CreateContainer(image)
	if err != nil {
		return nil, err
	}

	return &resp.ID, StartContainer(resp.ID)
}

func CreateContainer(image string) (container.CreateResponse, error) {
	resp, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: image,
	}, nil, nil, nil, "")
	return resp, err
}

func StartContainer(id string) error {
	return cli.ContainerStart(context.Background(), id, container.StartOptions{})
}

func StopContainer(id string) error {
	return cli.ContainerStop(context.Background(), id, container.StopOptions{})
}

func RemoveContainer(id string) error {
	return cli.ContainerRemove(context.Background(), id, container.RemoveOptions{})
}
