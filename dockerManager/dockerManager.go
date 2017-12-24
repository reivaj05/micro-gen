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

func (client *dockerClient) getImages(serviceName string) []*GoJSON.JSONWrapper {

	return nil
}

// func (client *dockerClient) filterImages() {

// }
