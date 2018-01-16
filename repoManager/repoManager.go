package repoManager

import (
	"fmt"
	"os"
	"os/exec"
)

type repoProviderClient interface {
	CreateCloudRepo(string) (string, error)
}

var githubKey = "GITHUB_TOKEN"
var bitbucketKey = "BITBUCKET_TOKEN"
var gitlabKey = "GITLAB_TOKEN"

var repoProviderKeys = map[string]string{
	"github":    githubKey,
	"bitbucket": bitbucketKey,
	"gitlab":    gitlabKey,
}

var repoProviderClients = map[string]func(string) repoProviderClient{
	"github": NewGithubClient,
	// "bitbucket": NewBitbucketClient,
	"gitlab": NewGitlabClient,
}

func CreateRepo(serviceName, provider string) error {
	if err := createLocalRepo(serviceName); err != nil {
		return err
	}
	return createRepoProvider(serviceName, provider)
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

func createRepoProvider(serviceName, provider string) error {
	token, err := getToken(provider)
	if err != nil {
		return err
	}
	client := repoProviderClients[provider](token)
	fmt.Printf("Creating %s repository in %s provider...\n", serviceName, provider)
	repoURL, err := client.CreateCloudRepo(serviceName)
	if err != nil {
		return err
	}
	return linkCloudRepoToLocalRepo(repoURL, serviceName)
}

func getToken(provider string) (string, error) {
	if key, ok := repoProviderKeys[provider]; ok {
		accessToken := os.Getenv(key)
		if accessToken == "" {
			return "", fmt.Errorf("%s env var does not exist", key)
		}
		return accessToken, nil
	}
	return "", fmt.Errorf("Repo provider '%s' not supported", provider)
}

func linkCloudRepoToLocalRepo(repoURL, serviceName string) error {
	fmt.Printf("Linking %s repository to local repository...\n", serviceName)
	cmd := exec.Command("sh", "-c", fmt.Sprintf("cd %s && git remote add origin %s", serviceName, repoURL))
	_, err := cmd.CombinedOutput()
	return err
}
