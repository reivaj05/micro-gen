package rustGenerator

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/reivaj05/GoConfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RustBuilderTestSuite struct {
	suite.Suite
	assert      *assert.Assertions
	pathBackup  string
	serviceName string
}

func (suite *RustBuilderTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.serviceName = "mockService"
	GoConfig.Init(&GoConfig.ConfigOptions{
		ConfigType: "json",
		ConfigFile: "config",
		ConfigPath: "../../../",
	})
}

func (suite *RustBuilderTestSuite) TearDownSuite() {
	os.RemoveAll(fmt.Sprintf("./%s", suite.serviceName))
}

func (suite *RustBuilderTestSuite) TestBuildWrongServiceName() {
	serviceName := strings.Repeat("f", 1000)
	err := Build(serviceName)
	suite.assert.NotNil(err)
	_, err = os.Stat(fmt.Sprintf("./%s", serviceName))
	suite.assert.False(os.IsNotExist(err))
}

func (suite *RustBuilderTestSuite) TestBuildWrongTemplatesPath() {
	GoConfig.SetConfigValue("templates", "wrongPath")
	err := Build(suite.serviceName)
	suite.assert.NotNil(err)
}

func (suite *RustBuilderTestSuite) TestBuildSuccessful() {
	err := Build(suite.serviceName)
	suite.assert.Nil(err)
}

func TestRustBuilder(t *testing.T) {
	suite.Run(t, new(RustBuilderTestSuite))
}
