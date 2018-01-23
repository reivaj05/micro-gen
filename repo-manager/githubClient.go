package repoManager

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type githubClient struct {
	token string
}

func NewGithubClient(token string) repoProviderClient {
	return &githubClient{
		token: token,
	}
}

func (client *githubClient) CreateCloudRepo(serviceName string) (string, error) {
	ctx := context.Background()
	gClient := client.createGitHubClient(ctx)
	return client.createGithubRepo(serviceName, gClient, ctx)
}

func (client *githubClient) createGitHubClient(ctx context.Context) *github.Client {
	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: client.token}))
	return github.NewClient(tc)
}

func (client *githubClient) createGithubRepo(serviceName string, gClient *github.Client,
	ctx context.Context) (repoURL string, err error) {

	repo := &github.Repository{Name: github.String(serviceName)}
	repo, _, err = gClient.Repositories.Create(ctx, "", repo)
	if err != nil {
		return "", err
	}
	return *repo.SSHURL, nil
}
