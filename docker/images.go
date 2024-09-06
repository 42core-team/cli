package docker

import (
	"context"

	"github.com/docker/docker/api/types/image"
)

func PullImage(imageUrl string) error {
	_, err := cli.ImagePull(context.Background(), imageUrl, image.PullOptions{})
	// io.Copy(os.Stdout, reader)
	return err
}
