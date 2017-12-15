package CIManager

import (
	"github.com/reivaj05/GoRequester"
)

type travisClient struct {
	requesterObj *requester.Requester
	token        string
	headers      map[string]string
}

func New() *travisClient {
	return &travisClient{
	// TODO:
	}
}

func (client *travisClient) enableRepo(repoName string) (string, error) {
	// TODO:
	return "", nil
}

func (client *travisClient) filterReposByName(repoName string) []string {
	// TODO:
	return []string{}
}

func (client *travisClient) getRepo() (string, error) {
	// TODO:
	return "", nil
}

func (client *travisClient) getRepos() (string, error) {
	// TODO:
	return "", nil
}

func (client *travisClient) syncAccount() (string, error) {
	// TODO:
	return "", nil
}
