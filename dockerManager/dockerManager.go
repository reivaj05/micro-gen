package dockerManager

import (
	"github.com/reivaj05/GoJSON"
	"github.com/reivaj05/GoRequester"
)

type dockerClient struct {
	requesterObj *requester.Requester
	token        string
}

func NewDockerClient(token string) *dockerClient {
	return &dockerClient{
		requesterObj: requester.New(),
		token:        token,
	}
}

func (client *dockerClient) getRelatedImages(serviceName string) []string {
	// TODO: Implement
	return []string{}
}

func (client *dockerClient) filterImagesByName(serviceName string) []*GoJSON.JSONWrapper {
	// TODO: Implement
}

func (client *dockerClient) getImages() []*GoJSON.JSONWrapper {
	// TODO: Implement
}
