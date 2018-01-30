package toolingBuilder

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/reivaj05/GoConfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ManagerBuilderTestSuite struct {
	suite.Suite
	assert      *assert.Assertions
	pathBackup  string
	managerName string
	services    string
}

func (suite *ManagerBuilderTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.managerName = "mockManager"
	suite.services = "mockService1,mockService2,mockService3"
	GoConfig.Init(&GoConfig.ConfigOptions{
		ConfigType: "json",
		ConfigFile: "config",
		ConfigPath: "../../../",
	})
}

func (suite *ManagerBuilderTestSuite) TearDownSuite() {
	os.RemoveAll(fmt.Sprintf("./%s", suite.managerName))
}

func (suite *ManagerBuilderTestSuite) TestBuildWrongServiceName() {
	managerName := strings.Repeat("f", 1000)
	err := Build(managerName, suite.services)
	suite.assert.NotNil(err)
	_, err = os.Stat(fmt.Sprintf("./%s", managerName))
	suite.assert.False(os.IsNotExist(err))
}

func (suite *ManagerBuilderTestSuite) TestBuildWrongTemplatesPath() {
	GoConfig.SetConfigValue("templates", "wrongPath")
	err := Build(suite.managerName, suite.services)
	suite.assert.NotNil(err)
}

func (suite *ManagerBuilderTestSuite) TestBuildSuccessful() {
	err := Build(suite.managerName, suite.services)
	suite.assert.Nil(err)
}

func TestGoBuilder(t *testing.T) {
	suite.Run(t, new(ManagerBuilderTestSuite))
}
