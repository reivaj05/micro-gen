package CIManager

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type CIConnector func(string) error


var travisCIKey = "TRAVIS_TOKEN"
var jenkinsCIKey = "JENKINS_TOKEN"
var circleCIKey = "CIRCLE_TOKEN"

var CIConnectors = map[string]CIConnector{
	"travis":	travisConnector,
	"jenkins":  jenkinsConnector,
	"circle":   circleConnector,
}


func ConnectWithCIProvider(provider string) error {
	if creator, ok := repoProviders[provider]; ok {
		return creator(serviceName)
	}
	return fmt.Errorf("Repo provider '%s' not supported", provider)
}

func createLocalRepo(serviceName string) error {
	fmt.Println("Creating local repo...")
	cmd := exec.Command("sh", "-c", fmt.Sprintf("cd %s && git init", serviceName))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func createGithubRepo(serviceName string) error {
	fmt.Printf("Creating %s github repository...\n", serviceName)
	ctx := context.Background()
	client, err := createGitHubClient(ctx)
	if err != nil {
		return err
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
	ctx context.Context) (err error) {

	repo := &github.Repository{Name:    github.String(serviceName)}
	repo, _, err  = client.Repositories.Create(ctx, "", repo)
	if err != nil {
		return err
	}
	return linkGithubRepoToLocalRepo(repo, serviceName)
}

func linkGithubRepoToLocalRepo(repo *github.Repository, serviceName string) error {
	fmt.Printf("Linking %s repository to local repository...\n", serviceName)
	cmd := exec.Command("sh", "-c", fmt.Sprintf("cd %s && git remote add origin %s", serviceName, *repo.SSHURL))
	_, err := cmd.CombinedOutput()
	return err
}

func createBitbucketRepo(serviceName string) error {
	fmt.Println("TODO: Create bitbucket repo")
	return nil
}

func createGitlabRepo(serviceName string) error {
	fmt.Println("TODO: Create gitlab repo")
	return nil
}