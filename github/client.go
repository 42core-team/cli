package github

import (
	"io"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func RepoExists(path string) bool {
	_, err := git.PlainOpen(path)
	return err == nil
}

func Clone(url, path string) (*git.Repository, error) {
	return git.PlainClone(path, false, &git.CloneOptions{
		URL:          url,
		SingleBranch: true,
		Depth:        1,
		Progress:     io.Discard,
		Auth: &http.BasicAuth{
			Username: "nil",
			Password: os.Getenv("GITHUB_TOKEN"),
		},
	})
}

func Pull(path string) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	// Get the worktree of the repository
	w, err := repo.Worktree()
	if err != nil {
		return err
	}

	// Pull the latest changes
	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
		Progress:   io.Discard,
		Auth: &http.BasicAuth{
			Username: "nil",
			Password: os.Getenv("GITHUB_TOKEN"),
		},
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		return err
	}

	return nil
}

func CommitAndPush(path, message string) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	w, err := repo.Worktree()
	if err != nil {
		return err
	}

	_, err = w.Add(".")
	if err != nil {
		return err
	}

	// Complete the commit operation
	_, err = w.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "CORE CLI",        // Replace with your name
			Email: "cli@coregame.de", // Replace with your email
			When:  time.Now(),
		},
	})
	if err != nil {
		return err
	}

	// Push the changes using the default options
	err = repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: &http.BasicAuth{
			Username: "nil",                     // This can be anything except an empty string when using a token
			Password: os.Getenv("GITHUB_TOKEN"), // Assumes you have a GITHUB_TOKEN environment variable
		},
	})
	if err != nil {
		return err
	}

	return nil
}
