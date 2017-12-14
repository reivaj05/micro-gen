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

func ConnectWithCIProvider(serviceName, provider string) error {
	if connector, ok := CIConnectors[provider]; ok {
		return connector(serviceName)
	}
	return fmt.Errorf("CI provider '%s' not supported", provider)
}

func travisConnector(serviceName string) error {
	_, err := getToken(travisCIKey)
	if err != nil {
		return err
	}
	return nil
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
