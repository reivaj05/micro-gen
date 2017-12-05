package repoManager

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type repoCreator func(string) (*github.Repository, error)


var githubTokenKey = "GITHUB_TOKEN"
var bitbucketTokenKey = "BITBUCKET_TOKEN"
var gitlabTokenKey = "GITLAB_TOKEN"

var repoProviders = map[string]repoCreator{
	"github":		createGithubRepo,
	"bitbucket":    createBitbucketRepo,
	"gitlab":       createGitlabRepo,
}


func CreateRepo(serviceName, provider string) error {
	if err := createLocalRepo(serviceName); err != nil {
		return err
	}
	if creator, ok := repoProviders[provider]; ok {
		fmt.Printf("Creating %s repository...\n", provider)
		repo, err := creator(serviceName)
		if err != nil {
			return err
		}
		fmt.Printf("Linking %s repository to local repository...\n", serviceName)
		return linkGithubRepoToLocalRepo(repo)
	}
	return fmt.Errorf("Repo provider '%s' not supported", provider)
}

func createLocalRepo(serviceName string) error {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("cd %s; git init", serviceName))
	fmt.Println("Creating local repo...")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func createGithubRepo(serviceName string) (*github.Repository, error) {
	ctx := context.Background()
	client, err := createGitHubClient(ctx)
	if err != nil {
		return nil, err
	}
	return __createGithubRepo(serviceName, client, ctx)
}

func createGitHubClient(ctx context.Context) (*github.Client, error) {
	accessToken, err := getToken(githubTokenKey)
	if err != nil {
		return nil, err
	}
	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken}))
	return github.NewClient(tc), nil
}

func getToken(key string) (string, error) {
	accessToken := os.Getenv(key)
	if accessToken == "" {
		return "", fmt.Errorf("%s env var does not exist", key)
	}
	return accessToken, nil
}

func __createGithubRepo(serviceName string, client *github.Client,
	ctx context.Context) (repo *github.Repository, err error) {

	repo = &github.Repository{Name:    github.String(serviceName)}
	repo, _, err  = client.Repositories.Create(ctx, "", repo)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func createBitbucketRepo(serviceName string) (*github.Repository, error) {
	fmt.Println("TODO: Create bitbucket repo")
	return nil, nil
}

func createGitlabRepo(serviceName string) (*github.Repository, error) {
	fmt.Println("TODO: Create gitlab repo")
	return nil, nil
}

func linkGithubRepoToLocalRepo(repo *github.Repository) error {
	return nil
}