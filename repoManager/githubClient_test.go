package repoManager

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GithubClientTestSuite struct {
	suite.Suite
	assert      *assert.Assertions
	serviceName string
	token       string
}

func (suite *GithubClientTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.serviceName = "mockService"
	suite.token = "mockToken"
}

func (suite *GithubClientTestSuite) TestCreateGithubRepoSuccessfully() {
	client := NewGithubClient(suite.token)
	url, err := client.CreateCloudRepo(suite.serviceName)
	fmt.Println(url, err)
	// TODO: Find a way to mock successful behavior.
	// suite.assert.Nil(err)
	// suite.assert.NotEqual(url, "")
}

func (suite *GithubClientTestSuite) TestCreateGithubRepoUnsuccessfully() {
	client := NewGithubClient(suite.token)
	url, err := client.CreateCloudRepo(suite.serviceName)
	suite.assert.NotNil(err)
	suite.assert.Equal(url, "")
}

func TestGithubClient(t *testing.T) {
	suite.Run(t, new(GithubClientTestSuite))
}
