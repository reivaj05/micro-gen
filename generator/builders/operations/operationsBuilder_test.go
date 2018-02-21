package operationsBuilder

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/reivaj05/GoConfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type OperationsBuilderTestSuite struct {
	suite.Suite
	assert         *assert.Assertions
	operationsName string
	services       string
}

func (suite *OperationsBuilderTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.operationsName = "mockOperations"
	suite.services = "mockOperations1,mockOperations2,mockOperations3"
	GoConfig.Init(&GoConfig.ConfigOptions{
		ConfigType: "json",
		ConfigFile: "config",
		ConfigPath: "../../../",
	})
	suite.clearDockerCredentials()
}

func (suite *OperationsBuilderTestSuite) clearDockerCredentials() {
	os.Setenv("DOCKER_USERNAME", "")
	os.Setenv("DOCKER_PASSWORD", "")
	os.Setenv("DOCKER_REGISTRY", "")
}

func (suite *OperationsBuilderTestSuite) TearDownSuite() {
	os.RemoveAll(fmt.Sprintf("./%s", suite.operationsName))
}

func (suite *OperationsBuilderTestSuite) TestBuildWrongOperationsName() {
	operationsName := strings.Repeat("f", 1000)
	err := Build(operationsName, suite.services)
	suite.assert.NotNil(err)
	_, err = os.Stat(fmt.Sprintf("./%s", operationsName))
	suite.assert.False(os.IsNotExist(err))
}

func (suite *OperationsBuilderTestSuite) TestBuildWrongTemplatesPath() {
	GoConfig.SetConfigValue("templates", "wrongPath")
	err := Build(suite.operationsName, suite.services)
	suite.assert.NotNil(err)
}

func (suite *OperationsBuilderTestSuite) TestBuildSuccessful() {
	err := Build(suite.operationsName, suite.services)
	suite.assert.Nil(err)
}

func TestOperationsBuilder(t *testing.T) {
	suite.Run(t, new(OperationsBuilderTestSuite))
}
