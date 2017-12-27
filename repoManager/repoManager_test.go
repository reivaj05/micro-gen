package repoManager

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RepoManagerTestSuite struct {
	suite.Suite
	assert            *assert.Assertions
	managerName       string
	githubProvider    string
	gitlabProvider    string
	bitbucketProvider string
	githubToken       string
}

func (suite *RepoManagerTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.managerName = "mockManager"
	suite.githubProvider = "github"
	suite.gitlabProvider = "gitlab"
	suite.bitbucketProvider = "bitbucket"
	os.MkdirAll(suite.managerName, os.ModePerm)
}

func (suite *RepoManagerTestSuite) SetupTest() {
	suite.githubToken = os.Getenv(githubTokenKey)
}

func (suite *RepoManagerTestSuite) TearDownSuite() {
	os.RemoveAll(suite.managerName)
}

func (suite *RepoManagerTestSuite) TearDownTest() {
	os.Setenv(githubTokenKey, suite.githubToken)
}

func (suite *RepoManagerTestSuite) TestCreateGithubRepoSuccessfully() {
	os.Setenv(githubTokenKey, "GITHUB_MOCK_TOKEN")
	err := CreateRepo(suite.managerName, suite.githubProvider)
	suite.assert.Nil(err)
}

func (suite *RepoManagerTestSuite) TestCreateGithubRepoWithoutAccessToken() {
	os.Setenv(githubTokenKey, "")
	err := CreateRepo(suite.managerName, suite.githubProvider)
	suite.assert.NotNil(err)
}

func (suite *RepoManagerTestSuite) TestCreateGithubRepoWithWrongAccessToken() {
	os.Setenv(githubTokenKey, "GITHUB_MOCK_TOKEN")
	err := CreateRepo(suite.managerName, suite.githubProvider)
	suite.assert.NotNil(err)
}

func (suite *RepoManagerTestSuite) TestCreateGitlabRepoSuccessfully() {
	err := CreateRepo(suite.managerName, suite.gitlabProvider)
	suite.assert.Nil(err)
}

func (suite *RepoManagerTestSuite) TestCreateBitbucketRepoSuccessfully() {
	err := CreateRepo(suite.managerName, suite.bitbucketProvider)
	suite.assert.Nil(err)
}

func (suite *RepoManagerTestSuite) TestCreateRepoWithLocalRepoError() {
	// TODO: TestCreateRepoUnsuccessfulLocalRepo
}

func (suite *RepoManagerTestSuite) TestCreateRepoWithWrongServiceProvider() {
	err := CreateRepo(suite.managerName, "wrongServiceProvider")
	suite.assert.NotNil(err)
}

func TestRepoManager(t *testing.T) {
	suite.Run(t, new(RepoManagerTestSuite))
}
