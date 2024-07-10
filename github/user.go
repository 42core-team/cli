package github

import "github.com/google/go-github/v62/github"

func GetGithubUserByUserName(userName string) (*github.User, error) {
	user, _, err := client.Users.Get(getGithubContext(), userName)
	return user, err
}

func GithubUserExists(userName string) bool {
	_, err := GetGithubUserByUserName(userName)
	return err == nil
}
