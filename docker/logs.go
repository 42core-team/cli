package docker

import (
	"bytes"
	"context"
	"io"

	"github.com/docker/docker/api/types/container"
)

func GetLogs(id string) (string, error) {
	out, err := cli.ContainerLogs(context.Background(), id, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		return "", err
	}
	defer out.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, out)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
