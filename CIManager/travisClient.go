package CIManager

import (
	"fmt"

	"github.com/reivaj05/GoRequester"
)

type travisClient struct {
	requesterObj *requester.Requester
	token        string
	headers      map[string]string
}

func NewTravisClient(token string) *travisClient {
	return &travisClient{
		requesterObj: requester.New(),
		token:        token,
		headers:      createTravisRequestHeaders(token),
	}
}

func createTravisRequestHeaders(token string) map[string]string {
	return map[string]string{
		"Travis-API-Version": "3",
		"Authorization":      fmt.Sprintf("token %s", token),
	}
}

func (client *travisClient) ActivateRepo(serviceName string) error {
	// TODO:
	fmt.Println("Activate repo for ", serviceName)
	return nil
}

func (client *travisClient) filterReposByName(serviceName string) []string {
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
