package goBuilder

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/reivaj05/GoConfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GoBuilderTestSuite struct {
	suite.Suite
	assert     *assert.Assertions
	pathBackup string
}

func (suite *GoBuilderTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	GoConfig.Init(&GoConfig.ConfigOptions{
		ConfigType: "json",
		ConfigFile: "config",
		ConfigPath: "..",
	})
}

func (suite *GoBuilderTestSuite) SetupTest() {
	suite.pathBackup = GoConfig.GetConfigStringValue("goTemplatesPath")
}

func (suite *GoBuilderTestSuite) TearDownTest() {
	GoConfig.SetConfigValue("goTemplatesPath", suite.pathBackup)
}

func (suite *GoBuilderTestSuite) TestBuildWrongServiceName() {
	serviceName := strings.Repeat("f", 1000)
	err := Build(serviceName)
	suite.assert.NotNil(err)
	_, err = os.Stat(fmt.Sprintf("./%s", serviceName))
	suite.assert.False(os.IsNotExist(err))
}

func TestGoBuilder(t *testing.T) {
	suite.Run(t, new(GoBuilderTestSuite))
}
