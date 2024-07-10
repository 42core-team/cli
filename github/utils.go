package github

import (
	"context"

	"github.com/google/go-github/v62/github"
)

func getGithubContext() context.Context {
	return context.WithValue(context.Background(), github.SleepUntilPrimaryRateLimitResetWhenRateLimited, true)
}
