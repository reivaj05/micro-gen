package CIManager

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type CIConnector func(string) error


var travisCIKey = "TRAVIS_TOKEN"
var jenkinsCIKey = "JENKINS_TOKEN"
var circleCIKey = "CIRCLE_TOKEN"

var CIConnectors = map[string]CIConnector{
	"travis":	travisConnector,
	"jenkins":  jenkinsConnector,
	"circle":   circleConnector,
}


func ConnectWithCIProvider(serviceName, provider string) error {
	if connector, ok := CIConnectors[provider]; ok {
		return connector(serviceName)
	}
	return fmt.Errorf("CI provider '%s' not supported", provider)
}

func connectTravis(serviceName string) error {
	accessToken, err := getToken(travisCIKey)
	if err != nil {
		return nil, err
	}
	return nil
}

func connectJenkins(serviceName string) error {
	accessToken, err := getToken(jenkinsCIKey)
	if err != nil {
		return nil, err
	}
	return nil
}

func connectCircle(serviceName string) error {
	accessToken, err := getToken(circleCIKey)
	if err != nil {
		return nil, err
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