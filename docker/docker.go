package docker

import "github.com/docker/docker/client"

var cli *client.Client

func NewDockerClient() error {
	var err error
	cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	return err
}

func CloseDockerClient() error {
	return cli.Close()
}
