package repoManager

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var successfulToken = "SUCCESSFUL_TOKEN"
var localRepoLinkingError = "LINKING_ERROR_TOKEN"

type RepoManagerTestSuite struct {
	suite.Suite
	assert            *assert.Assertions
	managerName       string
	githubProvider    string
	gitlabProvider    string
	bitbucketProvider string
	githubToken       string
	bitbucketToken    string
	gitlabToken       string
}

type mockClient struct {
	token string
}

func (suite *RepoManagerTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.managerName = "mockManager"
	suite.githubProvider = "github"
	suite.gitlabProvider = "gitlab"
	suite.bitbucketProvider = "bitbucket"
	repoProviderClients = map[string]func(string) repoProviderClient{
		"github":    NewMockClient,
		"bitbucket": NewMockClient,
		"gitlab":    NewMockClient,
	}
}

func NewMockClient(token string) repoProviderClient {
	return &mockClient{token: token}
}

func (client *mockClient) CreateCloudRepo(serviceName string) (string, error) {
	if client.token == successfulToken {
		return "mockRemote", nil
	}
	if client.token == localRepoLinkingError {
		return "", nil
	}
	return "", fmt.Errorf("")
}

func (suite *RepoManagerTestSuite) SetupTest() {
	os.MkdirAll(suite.managerName, os.ModePerm)
	suite.githubToken = os.Getenv(githubKey)
	suite.bitbucketToken = os.Getenv(bitbucketKey)
	suite.gitlabToken = os.Getenv(gitlabKey)
}

func (suite *RepoManagerTestSuite) TearDownTest() {
	os.RemoveAll(suite.managerName)
	os.Setenv(githubKey, suite.githubToken)
	os.Setenv(bitbucketKey, suite.bitbucketToken)
	os.Setenv(gitlabKey, suite.gitlabToken)
}

func (suite *RepoManagerTestSuite) TestCreateGithubRepoSuccessfully() {
	os.Setenv(githubKey, successfulToken)
	err := CreateRepo(suite.managerName, suite.githubProvider)
	suite.assert.Nil(err)
}

func (suite *RepoManagerTestSuite) TestCreateGithubRepoUnsuccessfully() {
	err := CreateRepo(suite.managerName, suite.githubProvider)
	suite.assert.NotNil(err)
}

func (suite *RepoManagerTestSuite) TestCreateBitbucketRepoSuccessfully() {
	os.Setenv(bitbucketKey, successfulToken)
	err := CreateRepo(suite.managerName, suite.bitbucketProvider)
	suite.assert.Nil(err)
}

func (suite *RepoManagerTestSuite) TestCreateBitbucketRepoUnsuccessfully() {
	err := CreateRepo(suite.managerName, suite.bitbucketProvider)
	suite.assert.NotNil(err)
}

func (suite *RepoManagerTestSuite) TestCreateGitlabRepoSuccessfully() {
	os.Setenv(gitlabKey, successfulToken)
	err := CreateRepo(suite.managerName, suite.gitlabProvider)
	suite.assert.Nil(err)
}

func (suite *RepoManagerTestSuite) TestCreateGitlabRepoUnsuccessfully() {
	err := CreateRepo(suite.managerName, suite.gitlabProvider)
	suite.assert.NotNil(err)
}

func (suite *RepoManagerTestSuite) TestCreateRepoWithLocalRepoError() {
	tmp := suite.managerName
	suite.managerName = "wrongDir"
	err := CreateRepo(suite.managerName, suite.githubProvider)
	suite.assert.NotNil(err)
	suite.managerName = tmp
}

func (suite *RepoManagerTestSuite) TestCreateRepoWithLinkingLocalRepoError() {
	os.Setenv(githubKey, localRepoLinkingError)
	err := CreateRepo(suite.managerName, suite.githubProvider)
	suite.assert.NotNil(err)
}

func (suite *RepoManagerTestSuite) TestCreateRepoWithWrongServiceProvider() {
	err := CreateRepo(suite.managerName, "wrongRepoProvider")
	suite.assert.NotNil(err)
}

func TestRepoManager(t *testing.T) {
	suite.Run(t, new(RepoManagerTestSuite))
}
