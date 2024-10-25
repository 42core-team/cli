package github

import (
	"context"
	"core-cli/utils"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"
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
	r := &github.TemplateRepoRequest{
		Name:    github.String(name),
		Owner:   github.String(orgName),
		Private: github.Bool(true),
	}

	repo, _, err := client.Repositories.CreateFromTemplate(getGithubContext(), *template.Owner.Login, *template.Name, r)
	return repo, err
}

func CreateForkRepo(name string, fork *github.Repository) (*github.Repository, error) {
	r := &github.RepositoryCreateForkOptions{
		Name:         name,
		Organization: orgName,
	}

	repo, _, err := client.Repositories.CreateFork(getGithubContext(), *fork.Owner.Login, *fork.Name, r)
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

func ChangeCollaboratorReadOnly(repoName, userName string) error {
	_, _, err := client.Repositories.AddCollaborator(getGithubContext(), orgName, repoName, userName, &github.RepositoryAddCollaboratorOptions{
		Permission: "pull",
	})
	return err
}

func ChangeCollaboratorInviteReadOnly(repoName, userName string) error {
	invites, _, err := client.Repositories.ListInvitations(getGithubContext(), orgName, userName, nil)
	if err != nil {
		return err
	}

	for _, invite := range invites {
		if invite.GetPermissions() != "pull" {
			_, _, err := client.Repositories.UpdateInvitation(getGithubContext(), orgName, repoName, invite.GetID(), "read")
			if err != nil {
				log.Default().Println(err)
			}
		}
	}

	return nil
}

func CreateTraceRelease(repoName, body string, draft, prerelease bool) (*github.RepositoryRelease, error) {
	ctx := context.Background()

	// Function to check if a tag exists
	tagExists := func(tag string) (bool, error) {
		_, _, err := client.Repositories.GetReleaseByTag(ctx, orgName, repoName, tag)
		if err != nil {
			if _, ok := err.(*github.ErrorResponse); ok && err.(*github.ErrorResponse).Response.StatusCode == 404 {
				return false, nil
			}
			return false, err
		}
		return true, nil
	}

	// Increment tag if it already exists
	tagName := "trace00"
	for {
		exists, err := tagExists(tagName)
		if err != nil {
			return nil, err
		}
		if !exists {
			break
		}
		tagName = incrementTraceTag(tagName)
	}

	body = utils.TrimStringWithIndicator(body, 125000, "....[Log Truncated]....")

	release := &github.RepositoryRelease{
		TagName:    github.String(tagName),
		Name:       github.String(tagName),
		Body:       github.String(body),
		Draft:      github.Bool(draft),
		Prerelease: github.Bool(prerelease),
	}

	repoRelease, _, err := client.Repositories.CreateRelease(ctx, orgName, repoName, release)
	if err != nil {
		return nil, err
	}

	return repoRelease, nil
}

func incrementTraceTag(tag string) string {
	numStr := tag[5:]
	num, err := strconv.Atoi(numStr)
	if err == nil {
		return fmt.Sprintf("trace%02d", num+1)
	}
	return "trace00"
}
