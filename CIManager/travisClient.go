package CIManager

import (
	"fmt"

	"github.com/reivaj05/GoJSON"
	"github.com/reivaj05/GoRequester"
)

var reposEndpoint = "https://api.travis-ci.org/repos"

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
	_, err := client.filterRepoByName(serviceName)
	if err != nil {
		return err
	}
	return nil
}

func (client *travisClient) filterRepoByName(serviceName string) (string, error) {
	// TODO:
	fmt.Println("Filterrepo by name")
	repos, err := client.getRepos()
	fmt.Println(repos)
	if err != nil {
		return "", err
	}
	return "", nil
}

// func (client *travisClient) filterBy(key, query string, repos []string) (string, error) {
// 	for repo, _ := range repos {
// 		if key in repo {
// 			if repo[key] == query {
// 				return "", nil
// 			}
// 		}
// 	}
// 	return "", fmt.Errorf("The repo %s couldn't be found", query)
// }

func (client *travisClient) getRepos() ([]*GoJSON.JSONWrapper, error) {
	jsonResponse, err := client.__getReposRequest()
	if err != nil {
		return nil, err
	}
	return jsonResponse.GetArrayFromPath("repositories"), nil
}

func (client *travisClient) getRepo() (string, error) {
	// TODO:
	return "", nil
}

func (client *travisClient) __getReposRequest() (*GoJSON.JSONWrapper, error) {
	config := client.createTravisRequestConfig("GET", reposEndpoint)
	response, _, err := client.requesterObj.MakeRequest(config)
	if err != nil {
		return nil, err
	}
	return GoJSON.New(response)
}

func (client *travisClient) syncAccount() (string, error) {
	// TODO:
	return "", nil
}

func (client *travisClient) createTravisRequestConfig(method, url string) *requester.RequestConfig {
	return &requester.RequestConfig{
		Method:  method,
		URL:     url,
		Headers: client.headers,
	}
}
