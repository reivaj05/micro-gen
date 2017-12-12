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
}

func (suite *RepoManagerTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.managerName = "mockManager"
	suite.githubProvider = "github"
	suite.gitlabProvider = "gitlab"
	suite.bitbucketProvider = "bitbucket"
}

func (suite *RepoManagerTestSuite) TearDownSuite() {
	// TODO: Add teardown
}

func (suite *RepoManagerTestSuite) TestCreateGithubRepoSuccessfully() {
	originalToken := os.Getenv("GITHUB_TOKEN")
	os.Setenv("GITHUB_TOKEN", "GITHUB_MOCK_TOKEN")
	err := CreateRepo(suite.managerName, suite.githubProvider)
	suite.assert.NotNil(err)
	os.Setenv("GITHUB_TOKEN", originalToken)
}

func (suite *RepoManagerTestSuite) TestCreateGithubRepoWithoutAccessToken() {
	// TODO: TestCreateGithubRepoWithoutAccessToken
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
