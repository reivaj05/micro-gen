package dockerManager

import requester "github.com/reivaj05/GoRequester"

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
