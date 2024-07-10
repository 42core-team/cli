package github

import (
	"github.com/google/go-github/v62/github"
)

func CreateRepo(name string) (*github.Repository, error) {
	r := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(true),
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
