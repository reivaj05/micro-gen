package repoManager

import "fmt"

type bitbucketClient struct {
	token string
}

func NewBitbucketClient(token string) repoProviderClient {
	return &bitbucketClient{
		token: token,
	}
}

func (client *bitbucketClient) CreateCloudRepo(serviceName string) (string, error) {
	fmt.Println("TODO: Implement bitbucket client")
	return "", nil
}
