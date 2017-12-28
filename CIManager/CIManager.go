package CIManager

import (
	"fmt"
	"os"
)

type CIConnector func(string, CIClient) error

type CIClient interface {
	ActivateRepo(string) error
}

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

var CIClients = map[string]func(string) CIClient{
	"travis": NewTravisClient,
}

func ConnectWithCIProvider(serviceName, provider string) error {
	token, err := getToken(provider)
	if err != nil {
		return err
	}
	return CIConnectors[provider](serviceName, CIClients[provider](token))
}

func travisConnector(serviceName string, client CIClient) error {
	return client.ActivateRepo(serviceName)
}

func jenkinsConnector(serviceName string, client CIClient) error {
	return nil
}

func circleConnector(serviceName string, client CIClient) error {
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
