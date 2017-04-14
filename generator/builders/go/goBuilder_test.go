package goBuilder

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GoBuilderTestSuite struct {
	suite.Suite
	assert *assert.Assertions
}

func (suite *GoBuilderTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	// GoConfig.Init(&GoConfig.ConfigOptions{
	// 	ConfigType: "json",
	// 	ConfigFile: "config",
	// 	ConfigPath: "..",
	// })
}

func (suite *GoBuilderTestSuite) TearDownSuite() {
	// rollback(suite.mockServiceName)
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
