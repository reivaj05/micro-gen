package CIManager

import (
	"fmt"
	"os"

	"github.com/reivaj05/GoRequester"
)

type CIConnector func(string) error

var travisCIKey = "TRAVIS_TOKEN"
var jenkinsCIKey = "JENKINS_TOKEN"
var circleCIKey = "CIRCLE_TOKEN"

var requesterObj = requester.New()

var travisActivateEndpoint = "https://api.travis-ci.org/repos"

var CIConnectors = map[string]CIConnector{
	"travis":  travisConnector,
	"jenkins": jenkinsConnector,
	"circle":  circleConnector,
}

func ConnectWithCIProvider(serviceName, provider string) error {
	if connector, ok := CIConnectors[provider]; ok {
		return connector(serviceName)
	}
	return fmt.Errorf("CI provider '%s' not supported", provider)
}

func travisConnector(serviceName string) error {
	token, err := getToken(travisCIKey)
	if err != nil {
		return err
	}
	return activateRepoInTravis(token)
}

func activateRepoInTravis(token string) error {
	response, status_code, err := requesterObj.MakeRequest(createTravisRequestConfig(token))
	fmt.Println(response, status_code, err)
	return nil
}

func createTravisRequestConfig(token string) *requester.RequestConfig {
	return &requester.RequestConfig{
		Method:  "GET",
		URL:     travisActivateEndpoint,
		Headers: createTravisRequestHeaders(token),
	}
}

func createTravisRequestHeaders(token string) map[string]string {
	return map[string]string{
		"Travis-API-Version": "3",
		"Authorization":      fmt.Sprintf("token %s", token),
	}
}

func jenkinsConnector(serviceName string) error {
	_, err := getToken(jenkinsCIKey)
	if err != nil {
		return err
	}
	return nil
}

func circleConnector(serviceName string) error {
	_, err := getToken(circleCIKey)
	if err != nil {
		return err
	}
	return nil
}

func getToken(key string) (string, error) {
	accessToken := os.Getenv(key)
	if accessToken == "" {
		return "", fmt.Errorf("%s env var does not exist", key)
	}
	return accessToken, nil
}
