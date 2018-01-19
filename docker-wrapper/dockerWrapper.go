package dockerWrapper

import (
	"github.com/reivaj05/GoJSON"
)

func (client *dockerClient) getRelatedImages(serviceName string) []string {
	// TODO: Implement
	return []string{}
}

func (client *dockerClient) filterImagesByName(serviceName string) []*GoJSON.JSONWrapper {
	// TODO: Implement
	return nil
}

func (client *dockerClient) getImages() []*GoJSON.JSONWrapper {
	// TODO: Implement
	return nil
}
