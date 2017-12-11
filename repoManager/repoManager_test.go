package repoManager

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/reivaj05/GoConfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RepoManagerTestSuite struct {
	suite.Suite
	assert      *assert.Assertions
}

func (suite *RepoManagerTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.managerName = "mockManager"
}

func (suite *RepoManagerTestSuite) TearDownSuite() {
	// TODO: Add teardown
}

func (suite *RepoManagerTestSuite) TestCreateGithubRepoSuccessfully() {
	// TODO: TestCreateGithubRepoSuccessfully
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
	err := Build(suite.serviceName)
	suite.assert.Nil(err)
}

func TestRepoManager(t *testing.T) {
	suite.Run(t, new(RepoManagerTestSuite))
}
