package CIManager

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/reivaj05/GoJSON"
	"github.com/reivaj05/GoRequester"
)

const (
	MAX_ATTEMPTS = 3
)

var baseURL = "https://api.travis-ci.org"
var reposEndpoint = fmt.Sprintf("%s/repos", baseURL)
var repoActivateEndpoint = "%s/repo/%s/activate"
var userEndpoint = fmt.Sprintf("%s/user", baseURL)
var syncAccountEndpoint = "%s/user/%s/sync"
var envVarsEndpoint = "%s/repo/%s/env_vars"

type travisClient struct {
	requesterObj *requester.Requester
	token        string
	headers      map[string]string
	attempts     int
}

func NewTravisClient(token string) CIClient {
	return &travisClient{
		requesterObj: requester.New(),
		token:        token,
		headers:      createTravisRequestHeaders(token),
		attempts:     0,
	}
}

func createTravisRequestHeaders(token string) map[string]string {
	return map[string]string{
		"Travis-API-Version": "3",
		"Authorization":      fmt.Sprintf("token %s", token),
	}
}

func (client *travisClient) ActivateRepo(serviceName string) error {
	if err := client.syncAccount(); err != nil {
		return err
	}
	time.Sleep(5 * time.Second) // Wait 5 seconds until sync is done
	repo, err := client.filterRepoByName(serviceName)
	if err != nil {
		if client.attempts < MAX_ATTEMPTS {
			client.attempts++
			return client.ActivateRepo(serviceName)
		}
		return err
	}
	slug, _ := repo.GetStringFromPath("slug")
	client.createEnvVars(slug)
	return client.activateRepoRequest(slug)
}

func (client *travisClient) syncAccount() error {
	user, err := client.getCurrentUser()
	if err != nil {
		return err
	}
	id, _ := user.GetFloatFromPath("id")
	return client.syncAccountRequest(strconv.Itoa(int(id)))
}

func (client *travisClient) getCurrentUser() (*GoJSON.JSONWrapper, error) {
	config := client.createTravisRequestConfig("GET", userEndpoint)
	response, status, err := client.requesterObj.MakeRequest(config)
	if err := client.checkResponse(status, err); err != nil {
		return nil, err
	}
	return GoJSON.New(response)
}

func (client *travisClient) syncAccountRequest(userID string) error {
	config := client.createTravisRequestConfig("POST", fmt.Sprintf(syncAccountEndpoint, baseURL, userID))
	_, status, err := client.requesterObj.MakeRequest(config)
	return client.checkResponse(status, err)
}

func (client *travisClient) filterRepoByName(serviceName string) (*GoJSON.JSONWrapper, error) {
	repos, err := client.getRepos()
	if err != nil {
		return nil, err
	}
	return client.filterBy("name", serviceName, repos)
}

func (client *travisClient) getRepos() ([]*GoJSON.JSONWrapper, error) {
	jsonResponse, err := client.getReposRequest()
	if err != nil {
		return nil, err
	}
	return jsonResponse.GetArrayFromPath("repositories"), nil
}

func (client *travisClient) getReposRequest() (*GoJSON.JSONWrapper, error) {
	config := client.createTravisRequestConfig("GET", reposEndpoint)
	response, status, err := client.requesterObj.MakeRequest(config)
	if err := client.checkResponse(status, err); err != nil {
		return nil, err
	}
	return GoJSON.New(response)
}

func (client *travisClient) filterBy(key, query string,
	repos []*GoJSON.JSONWrapper) (*GoJSON.JSONWrapper, error) {
	for _, repo := range repos {
		if repo.HasPath(key) {
			if value, ok := repo.GetStringFromPath(key); ok && value == query {
				return repo, nil
			}
		}
	}
	return nil, fmt.Errorf("The repo %s couldn't be found", query)
}

func (client *travisClient) activateRepoRequest(repoID string) error {
	repoID = url.QueryEscape(repoID)
	config := client.createTravisRequestConfig("POST", fmt.Sprintf(repoActivateEndpoint, baseURL, repoID))
	_, _, err := client.requesterObj.MakeRequest(config)
	return err
}

func (client *travisClient) createEnvVars(repoID string) {
	repoID = url.QueryEscape(repoID)
	for _, key := range []string{"DOCKER_USERNAME", "DOCKER_PASSWORD"} {
		client.createEnvVar(key, repoID)
	}
}

func (client *travisClient) createEnvVar(key, repoID string) {
	if value := os.Getenv(key); value != "" {
		config := client.createTravisRequestConfig("POST", fmt.Sprintf(envVarsEndpoint, baseURL, repoID))
		config.Body = client.createEnvVarsBodyRequest(key)
		config.Headers["Content-Type"] = "application/json"
		client.requesterObj.MakeRequest(config)
	} else {
		fmt.Println(fmt.Sprintf("%s en var not set or it is empty, it will not be set in travis", key))
	}
}

func (client *travisClient) createEnvVarsBodyRequest(key string) []byte {
	data := `{
		"env_var.name": "%s", "env_var.value": "%s", "env_var.public": false
	}`
	return []byte(fmt.Sprintf(data, key, os.Getenv(key)))
}

func (client *travisClient) checkResponse(status int, err error) error {
	if err != nil {
		return err
	}
	// TODO: Refactor, do not harcode and generalize.
	if status >= 400 {
		return fmt.Errorf("Got response with status %d", status)
	}
	return nil
}

func (client *travisClient) createTravisRequestConfig(method, url string) *requester.RequestConfig {
	return &requester.RequestConfig{
		Method:  method,
		URL:     url,
		Headers: client.headers,
	}
}
