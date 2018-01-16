package repoManager

import "fmt"

type gitlabClient struct {
	token string
}

func NewGitlabClient(token string) repoProviderClient {
	return &gitlabClient{
		token: token,
	}
}

func (client *gitlabClient) CreateCloudRepo(serviceName string) (string, error) {
	fmt.Println("TODO: Implement gitlab client")
	return "", nil
}
