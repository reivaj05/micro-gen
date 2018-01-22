package dockerWrapper

import (
	"fmt"
	"os"

	"github.com/reivaj05/GoRequester"
)

var dockerUsernameKey = "DOCKER_USERNAME"
var dockerPasswordKey = "DOCKER_PASSWORD"
var dockerRegistryHostKey = "DOCKER_REGISTRY_HOST"

type dockerRegistryManager struct {
	client   *requester.Requester
	token    string
	host     string
	username string
}

func NewDockerRegistryManager() (*dockerRegistryManager, error) {
	if err := checkDockerRegistryCredentials(); err != nil {
		return nil, err
	}
	return &dockerRegistryManager{
		token:    getToken(),
		host:     os.Getenv(dockerRegistryHostKey),
		username: os.Getenv(dockerUsernameKey),
		client:   requester.New(),
	}, nil
}

func checkDockerRegistryCredentials() error {
	if err := checkDockerCredentials(); err != nil {
		return err
	}
	if value := os.Getenv(dockerRegistryHostKey); value == "" {
		return fmt.Errorf("Env var %s not set", dockerRegistryHostKey)
	}
	return nil
}

func getToken() string {
	username, password := os.Getenv(dockerUsernameKey), os.Getenv(dockerPasswordKey)
	jsonString := fmt.Sprintf(`{"username":"%s","password":"%s"}`, username, password)
	return base64.URLEncoding.EncodeToString([]byte(jsonString))
}

func checkDockerCredentials() error {
	for _, key := range []string{dockerUsernameKey, dockerPasswordKey, dockerRegistryHostKey} {
		if value := os.Getenv(dockerUsernameKey); value == "" {
			return fmt.Errorf("Env var %s not set", key)
		}
	}
	return nil
}

func (manager *dockerRegistryManager) SearchRepositories() {
	response, status, err := manager.client.MakeRequest(&requester.RequestConfig{
		Method:  "GET",
		URL:     fmt.Sprintf("%s/v2/repositories/%s/", manager.host, manager.username),
		Headers: map[string]string{"Authorization": " " + manager.token},
	})
	fmt.Println(response, status, err)

}
