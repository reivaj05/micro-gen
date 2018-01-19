package dockerWrapper

import (
	"encoding/base64"
	"encoding/json"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

type dockerManager struct {
	cli   *client.Client
	ctx   context.Context
	token string
}

func NewDockerManager() (*dockerManager, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	token, err := getAuthToken(cli)
	if err != nil {
		return nil, err
	}
	return &dockerManager{cli: cli, ctx: context.Background(), token: token}, nil
}

func getAuthToken(cli *client.Client) (string, error) {
	authConfig := types.AuthConfig{
		Username: "username",
		Password: "password",
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

}

func (manager *dockerManager) test() {

}
