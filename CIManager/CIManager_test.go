package CIManager

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CIManagerTestSuite struct {
	suite.Suite
	assert          *assert.Assertions
	serviceName     string
	travisProvider  string
	jenkinsProvider string
	circleProvider  string
	travisToken     string
}

func (suite *CIManagerTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.serviceName = "mockService"
	suite.travisProvider = "travis"
	suite.jenkinsProvider = "jenkins"
	suite.circleProvider = "circle"
}

func (suite *CIManagerTestSuite) SetupTest() {
	suite.travisToken = os.Getenv(travisCIKey)
}

func (suite *CIManagerTestSuite) TearDownTest() {
	os.Setenv(travisCIKey, suite.travisToken)
}

func (suite *CIManagerTestSuite) TestConnectWithTravisProviderSuccessfully() {
	os.Setenv(travisCIKey, "TRAVIS_MOCK_TOKEN")
	err := ConnectWithCIProvider(suite.serviceName, suite.travisProvider)
	suite.assert.Nil(err)
}

func (suite *CIManagerTestSuite) TestConnectWithTravisProviderUnsuccessfully() {
	os.Setenv(travisCIKey, "TRAVIS_MOCK_TOKEN")
	err := ConnectWithCIProvider(suite.serviceName, suite.travisProvider)
	suite.assert.NotNil(err)
}

func (suite *CIManagerTestSuite) TestConnectWithTravisProviderWithouToken() {
	os.Setenv(travisCIKey, "")
	err := ConnectWithCIProvider(suite.serviceName, suite.travisProvider)
	suite.assert.NotNil(err)
}

func TestCIManager(t *testing.T) {
	suite.Run(t, new(CIManagerTestSuite))
}
