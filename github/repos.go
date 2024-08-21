package github

import (
	"context"
	"errors"
	"net/url"
	"strings"

	"github.com/google/go-github/v62/github"
)

func GetRepoFromName(name string) (*github.Repository, error) {
	repo, _, err := client.Repositories.Get(getGithubContext(), orgName, name)
	return repo, err
}

func GetRepoFromURL(urlStr string) (*github.Repository, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	pathSegments := strings.Split(strings.Trim(parsedURL.Path, "/"), "/")
	if len(pathSegments) < 2 {
		return nil, errors.New("invalid URL format")
	}
	owner := pathSegments[0]
	repoName := pathSegments[1]

	repo, _, err := client.Repositories.Get(context.Background(), owner, repoName)
	return repo, err
}

func CreateRepo(name string) (*github.Repository, error) {
	r := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(true),
	}
	repo, _, err := client.Repositories.Create(getGithubContext(), orgName, r)
	return repo, err
}

func CreateRepoFromTemplate(name string, template *github.Repository) (*github.Repository, error) {
	r := &github.Repository{
		Name:               github.String(name),
		Private:            github.Bool(true),
		TemplateRepository: template,
	}
	repo, _, err := client.Repositories.Create(getGithubContext(), orgName, r)
	return repo, err
}

func DeleteRepo(name string) error {
	_, err := client.Repositories.Delete(getGithubContext(), orgName, name)
	return err
}

func AddCollaborator(repoName, userName string) error {
	_, _, err := client.Repositories.AddCollaborator(getGithubContext(), orgName, repoName, userName, &github.RepositoryAddCollaboratorOptions{
		Permission: "push",
	})
	return err
}
