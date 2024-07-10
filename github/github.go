package github

import (
	"context"
	"os"

	"github.com/google/go-github/v62/github"
)

var client *github.Client
var orgName = os.Getenv("GITHUB_ORG")

func NewClient() error {
	client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_TOKEN"))
	_, _, err := client.Organizations.List(context.Background(), "paulic", nil)
	return err
}
