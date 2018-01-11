package repoManager

import (
	"context"
	"fmt"

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
	fmt.Printf("Creating %s github repository...\n", serviceName)
	ctx := context.Background()
	gClient, err := client.createGitHubClient(ctx)
	if err != nil {
		return "", err
	}
	return client.createGithubRepo(serviceName, gClient, ctx)
}

func (client *githubClient) createGitHubClient(ctx context.Context) (*github.Client, error) {
	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: client.token}))
	return github.NewClient(tc), nil
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
