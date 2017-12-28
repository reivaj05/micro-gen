package CIManager

import (
	"fmt"
	"os"
)

type CIConnector func(string) error

var travisCIKey = "TRAVIS_TOKEN"
var jenkinsCIKey = "JENKINS_TOKEN"
var circleCIKey = "CIRCLE_TOKEN"

var CIConnectors = map[string]CIConnector{
	"travis":  travisConnector,
	"jenkins": jenkinsConnector,
	"circle":  circleConnector,
}

var CIKeys = map[string]string{
	"travis":  travisCIKey,
	"jenkins": jenkinsCIKey,
	"circle":  circleCIKey,
}

func ConnectWithCIProvider(serviceName, provider string) error {
	_, err := getToken(provider)
	if err != nil {
		return err
	}
	return CIConnectors[provider](serviceName)
}

func travisConnector(serviceName string) error {

	return activateRepoInTravis(serviceName, "")
}

func activateRepoInTravis(serviceName, token string) error {
	client := NewTravisClient(token)
	return client.ActivateRepo(serviceName)
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

func getToken(provider string) (string, error) {
	if key, ok := CIKeys[provider]; ok {
		accessToken := os.Getenv(key)
		if accessToken == "" {
			return "", fmt.Errorf("%s env var does not exist", key)
		}
		return accessToken, nil
	}
	return "", fmt.Errorf("CI provider '%s' not supported", provider)
}
