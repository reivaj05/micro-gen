package CIManager

import (
	"fmt"
	"os"
)

type CIClient interface {
	ActivateRepo(string) error
}

var travisCIKey = "TRAVIS_TOKEN"
var jenkinsCIKey = "JENKINS_TOKEN"
var circleCIKey = "CIRCLE_TOKEN"

var CIKeys = map[string]string{
	"travis":  travisCIKey,
	"jenkins": jenkinsCIKey,
	"circle":  circleCIKey,
}

var CIClients = map[string]func(string) CIClient{
	"travis": NewTravisClient,
	// "jenkins": NewJenkinsClient,
	// "circle": NewCircleClient,
}

func ConnectWithCIProvider(serviceName, provider string) error {
	token, err := getToken(provider)
	if err != nil {
		return err
	}
	client := CIClients[provider](token)
	return client.ActivateRepo(serviceName)
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
