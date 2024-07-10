package github

import (
	"context"
	"os"

	"github.com/google/go-github/v62/github"
)

var client *github.Client
var orgName string

func NewClient() error {
	orgName = os.Getenv("GITHUB_ORG")

	client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_TOKEN"))

	_, _, err := client.Organizations.List(context.Background(), "paulicstudios", nil)

	return err
}
