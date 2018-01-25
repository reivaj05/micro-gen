package dockerWrapper

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DockerWrapperTestSuite struct {
	suite.Suite
	assert *assert.Assertions
}

func (suite *DockerWrapperTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
}

func (suite *DockerWrapperTestSuite) SetupTest() {
}

func (suite *DockerWrapperTestSuite) TearDownTest() {
}

// func (suite *DockerWrapperTestSuite) TestCreateGithubRepoSuccessfully() {
// 	os.Setenv(githubKey, successfulToken)
// 	err := CreateRepo(suite.managerName, suite.githubProvider)
// 	suite.assert.Nil(err)
// }

func TestDockerWrapper(t *testing.T) {
	suite.Run(t, new(DockerWrapperTestSuite))
}
