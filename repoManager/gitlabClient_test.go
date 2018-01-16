package repoManager

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GitlabClientTestSuite struct {
	suite.Suite
	assert      *assert.Assertions
	serviceName string
	token       string
}

func (suite *GitlabClientTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.serviceName = "mockService"
	suite.token = "mockToken"
}

func (suite *GitlabClientTestSuite) TestCreateGitlabRepoSuccessfully() {
	client := NewGitlabClient(suite.token)
	url, err := client.CreateCloudRepo(suite.serviceName)
	suite.assert.Nil(err)
	suite.assert.NotEqual(url, "")
}

func (suite *GitlabClientTestSuite) TestCreateGitlabRepoUnsuccessfully() {
	client := NewGitlabClient(suite.token)
	url, err := client.CreateCloudRepo(suite.serviceName)
	suite.assert.NotNil(err)
	suite.assert.Equal(url, "")
}

func TestGitlabClient(t *testing.T) {
	suite.Run(t, new(GitlabClientTestSuite))
}
