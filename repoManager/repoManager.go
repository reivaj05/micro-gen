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
	// "gitlab":    NewGitlabClient,
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

type githubClient struct {
	token string
}

func NewGithubClient(token string) repoProviderClient {
	return &githubClient{
		token: token,
	}
}

func (client *githubClient) CreateCloudRepo(serviceName string) (string, error) {
	return "", nil
}

// func createGithubRepo(serviceName string) error {
// 	fmt.Printf("Creating %s github repository...\n", serviceName)
// 	ctx := context.Background()
// 	client, err := createGitHubClient(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	return __createGithubRepo(serviceName, client, ctx)
// }

// func createGitHubClient(ctx context.Context) (*github.Client, error) {
// 	accessToken, err := getToken(githubKey)
// 	if err != nil {
// 		return nil, err
// 	}
// 	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken}))
// 	return github.NewClient(tc), nil
// }

// func getToken(key string) (string, error) {
// 	accessToken := os.Getenv(key)
// 	if accessToken == "" {
// 		return "", fmt.Errorf("%s env var does not exist", key)
// 	}
// 	return accessToken, nil
// }

// func __createGithubRepo(serviceName string, client *github.Client,
// 	ctx context.Context) (err error) {

// 	repo := &github.Repository{Name: github.String(serviceName)}
// 	repo, _, err = client.Repositories.Create(ctx, "", repo)
// 	if err != nil {
// 		return err
// 	}
// 	return linkGithubRepoToLocalRepo(repo, serviceName)
// }

// func linkGithubRepoToLocalRepo(repo *github.Repository, serviceName string) error {
// 	fmt.Printf("Linking %s repository to local repository...\n", serviceName)
// 	cmd := exec.Command("sh", "-c", fmt.Sprintf("cd %s && git remote add origin %s", serviceName, *repo.SSHURL))
// 	_, err := cmd.CombinedOutput()
// 	return err
// }

// func createBitbucketRepo(serviceName string) error {
// 	fmt.Println("TODO: Create bitbucket repo")
// 	return nil
// }

// func createGitlabRepo(serviceName string) error {
// 	fmt.Println("TODO: Create gitlab repo")
// 	return nil
// }
