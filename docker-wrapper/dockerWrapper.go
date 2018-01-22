package dockerWrapper

import (
	"encoding/base64"
	// "encoding/json"
	"fmt"
	// "io"
	"os"

	// "github.com/docker/docker/api/types"
	// "github.com/docker/docker/client"
	// "github.com/heroku/docker-registry-client/registry"
	// "github.com/reivaj05/GoJSON"
	"github.com/reivaj05/GoRequester"
	// "golang.org/x/net/context"
)

var dockerUsernameKey = "DOCKER_USERNAME"
var dockerPasswordKey = "DOCKER_PASSWORD"
var dockerRegistryHostKey = "DOCKER_REGISTRY_HOST"

// type dockerManager struct {
// 	cli   *client.Client
// 	ctx   context.Context
// 	token string
// }

type dockerRegistryManager struct {
	client   *requester.Requester
	token    string
	host     string
	username string
}

// func NewDockerManager() (*dockerManager, error) {
// 	cli, err := client.NewEnvClient()
// 	if err != nil {
// 		return nil, err
// 	}
// 	token, err := getAuthToken(cli)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &dockerManager{cli: cli, ctx: context.Background(), token: token}, nil
// }

// func getAuthToken(cli *client.Client) (string, error) {
// 	if err := checkDockerCredentials(); err != nil {
// 		return "", err
// 	}
// 	encodedJSON, err := getAuthEncoded()
// 	if err != nil {
// 		return "", err
// 	}
// 	return base64.URLEncoding.EncodeToString(encodedJSON), nil
// }

func checkDockerCredentials() error {
	for _, key := range []string{dockerUsernameKey, dockerPasswordKey} {
		if value := os.Getenv(dockerUsernameKey); value == "" {
			return fmt.Errorf("Env var %s not set", key)
		}
	}
	return nil
}

// func getAuthEncoded() ([]byte, error) {
// 	authConfig := types.AuthConfig{
// 		Username: os.Getenv(dockerUsernameKey),
// 		Password: os.Getenv(dockerPasswordKey),
// 	}
// 	return json.Marshal(authConfig)
// }

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

func getToken() string {
	username, password := os.Getenv(dockerUsernameKey), os.Getenv(dockerPasswordKey)
	jsonString := fmt.Sprintf(`{"username":"%s","password":"%s"}`, username, password)
	return base64.URLEncoding.EncodeToString([]byte(jsonString))
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

func (manager *dockerRegistryManager) SearchRepositories() {
	response, status, err := manager.client.MakeRequest(&requester.RequestConfig{
		Method:  "GET",
		URL:     fmt.Sprintf("%s/v2/repositories/%s/", manager.host, manager.username),
		Headers: map[string]string{"Authorization": " " + manager.token},
	})
	fmt.Println(response, status, err)

}
