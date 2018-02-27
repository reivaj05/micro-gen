package CDManager

import (
	"fmt"
	"os"
)

type CDClient interface {
	DeployService(string) error
}

var awsKey = "AWS_TOKEN"
var gceKey = "GCE_TOKEN"
var herokuKey = "HEROKU_TOKEN"

var CDKeys = map[string]string{
	"aws":    awsKey,
	"gce":    gceKey,
	"heroku": herokuKey,
}

var CDClients = map[string]func(string) CDClient{
// "aws": NewAWSClient,
}

func DeployServiceWithCDProvider(serviceName, provider string) error {
	token, err := getToken(provider)
	if err != nil {
		return err
	}
	client := CDClients[provider](token)
	fmt.Println(fmt.Sprintf("Activating %s in %s provider", serviceName, provider))
	return client.DeployService(serviceName)
}

func getToken(provider string) (string, error) {
	if key, ok := CDKeys[provider]; ok {
		accessToken := os.Getenv(key)
		if accessToken == "" {
			return "", fmt.Errorf("%s env var does not exist", key)
		}
		return accessToken, nil
	}
	return "", fmt.Errorf("CD provider '%s' not supported", provider)
}
