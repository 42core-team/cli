package docker

import (
	"context"
	"io"

	"github.com/docker/docker/api/types/image"
)

func PullImage(imageUrl string) error {
	reader, err := cli.ImagePull(context.Background(), imageUrl, image.PullOptions{})
	io.Copy(io.Discard, reader)
	return err
}
