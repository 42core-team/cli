package github

import (
	"context"
	"os"

	"github.com/google/go-github/v62/github"
)

var client *github.Client

func NewClient() error {
	client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_TOKEN"))
	_, _, err := client.Organizations.List(context.Background(), "paulic", nil)
	return err
}
