package repoManager

import (
	// "fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RepoManagerTestSuite struct {
	suite.Suite
	assert      		*assert.Assertions
	managerName 		string
	githubProvider 		string
	gitlabProvider 		string
	bitbucketProvider 	string
	originalToken 		string
}

func (suite *RepoManagerTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.managerName = "mockManager"
	suite.githubProvider = "github"
	suite.gitlabProvider = "gitlab"
	suite.bitbucketProvider = "bitbucket"
}

func (suite *RepoManagerTestSuite) SetupTest() {
	suite.originalToken = os.Getenv("GITHUB_TOKEN")
}

func (suite *RepoManagerTestSuite) TearDownSuite() {
	// TODO: Add teardown
}

func (suite *RepoManagerTestSuite) TearDownTest() {
	os.Setenv("GITHUB_TOKEN", suite.originalToken)
}

func (suite *RepoManagerTestSuite) TestCreateGithubRepoSuccessfully() {
	os.Setenv("GITHUB_TOKEN", "GITHUB_MOCK_TOKEN")
	err := CreateRepo(suite.managerName, suite.githubProvider)
	suite.assert.Nil(err)
}

func (suite *RepoManagerTestSuite) TestCreateGithubRepoWithoutAccessToken() {
	os.Setenv("GITHUB_TOKEN", "")
	err := CreateRepo(suite.managerName, suite.githubProvider)
	suite.assert.NotNil(err)
}

func (suite *RepoManagerTestSuite) TestCreateGithubRepoWithWrongAccessToken() {
	// TODO: TestCreateGithubRepoWithWrongAccessToken
}

func (suite *RepoManagerTestSuite) TestCreateGitlabRepoSuccessfully() {
	// TODO: TestCreateGitlabRepoSuccessfully
}

func (suite *RepoManagerTestSuite) TestCreateBitbucketRepoSuccessfully() {
	// TODO: TestCreateBitbucketRepoSuccessfully
}

func (suite *RepoManagerTestSuite) TestCreateRepoWithLocalRepoError() {
	// TODO: TestCreateRepoUnsuccessfulLocalRepo
}

func (suite *RepoManagerTestSuite) TestCreateRepoWithWrongServiceProvider() {
	// TODO: TestCreateRepoWithWrongServiceProvider
}

func TestRepoManager(t *testing.T) {
	suite.Run(t, new(RepoManagerTestSuite))
}
