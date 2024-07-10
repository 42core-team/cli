package github

import (
	"os"

	"github.com/google/go-github/v62/github"
)

var client *github.Client

func NewClient() {
	client = github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_TOKEN"))
}
