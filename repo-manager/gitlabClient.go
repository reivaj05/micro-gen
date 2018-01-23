package repoManager

import (
	"github.com/xanzy/go-gitlab"

	"fmt"
)

type gitlabClient struct {
	token string
}

func NewGitlabClient(token string) repoProviderClient {
	return &gitlabClient{
		token: token,
	}
}

func (client *gitlabClient) CreateCloudRepo(serviceName string) (string, error) {
	gClient := gitlab.NewClient(nil, client.token)
	options := client.createProjectOptions(serviceName)
	return client.createGitlabProject(gClient, options)
}

func (client *gitlabClient) createGitlabProject(
	gClient *gitlab.Client, options *gitlab.CreateProjectOptions) (string, error) {

	project, _, err := gClient.Projects.CreateProject(options)
	if err != nil {
		return "", err
	}
	return project.SSHURLToRepo, nil
}

func (client *gitlabClient) createProjectOptions(serviceName string) *gitlab.CreateProjectOptions {
	return &gitlab.CreateProjectOptions{
		Name:                 gitlab.String(serviceName),
		Description:          gitlab.String(fmt.Sprintf("Project %s created with micro-gen", serviceName)),
		MergeRequestsEnabled: gitlab.Bool(true),
		SnippetsEnabled:      gitlab.Bool(true),
		Visibility:           gitlab.Visibility(gitlab.PublicVisibility),
	}
}
