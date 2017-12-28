package CIManager

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var successfulToken = "SUCCESSFUL_TOKEN"

type CIManagerTestSuite struct {
	suite.Suite
	assert          *assert.Assertions
	serviceName     string
	travisProvider  string
	jenkinsProvider string
	circleProvider  string
	travisToken     string
	jenkinsToken    string
	circleToken     string
}

type mockClient struct {
	token string
}

func (suite *CIManagerTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.serviceName = "mockService"
	suite.travisProvider = "travis"
	suite.jenkinsProvider = "jenkins"
	suite.circleProvider = "circle"
	CIClients = map[string]func(string) CIClient{
		"travis":  NewMockClient,
		"jenkins": NewMockClient,
		"circle":  NewMockClient,
	}
}

func NewMockClient(token string) CIClient {
	return &mockClient{token: token}
}

func (client *mockClient) ActivateRepo(serviceName string) error {
	if client.token == successfulToken {
		return nil
	}
	return fmt.Errorf("")
}

func (suite *CIManagerTestSuite) SetupTest() {
	suite.travisToken = os.Getenv(travisCIKey)
	suite.jenkinsToken = os.Getenv(jenkinsCIKey)
	suite.circleToken = os.Getenv(circleCIKey)
}

func (suite *CIManagerTestSuite) TearDownTest() {
	os.Setenv(travisCIKey, suite.travisToken)
	os.Setenv(jenkinsCIKey, suite.jenkinsToken)
	os.Setenv(circleCIKey, suite.circleToken)
}

func (suite *CIManagerTestSuite) TestConnectWithTravisProviderSuccessfully() {
	os.Setenv(travisCIKey, successfulToken)
	err := ConnectWithCIProvider(suite.serviceName, suite.travisProvider)
	suite.assert.Nil(err)
}

func (suite *CIManagerTestSuite) TestConnectWithTravisProviderUnsuccessfully() {
	err := ConnectWithCIProvider(suite.serviceName, suite.travisProvider)
	suite.assert.NotNil(err)
}

func (suite *CIManagerTestSuite) TestConnectWithJenkinsProviderSuccessfully() {
	os.Setenv(jenkinsCIKey, successfulToken)
	err := ConnectWithCIProvider(suite.serviceName, suite.jenkinsProvider)
	suite.assert.Nil(err)
}

func (suite *CIManagerTestSuite) TestConnectWithJenkinsProviderUnsuccessfully() {
	err := ConnectWithCIProvider(suite.serviceName, suite.jenkinsProvider)
	suite.assert.NotNil(err)
}

func (suite *CIManagerTestSuite) TestConnectWithCircleProviderSuccessfully() {
	os.Setenv(circleCIKey, successfulToken)
	err := ConnectWithCIProvider(suite.serviceName, suite.circleProvider)
	suite.assert.Nil(err)
}

func (suite *CIManagerTestSuite) TestConnectWithCircleProviderUnsuccessfully() {
	err := ConnectWithCIProvider(suite.serviceName, suite.circleProvider)
	suite.assert.NotNil(err)
}

func (suite *CIManagerTestSuite) TestConnectWithCIProviderProviderWithoutToken() {
	os.Setenv(travisCIKey, "")
	err := ConnectWithCIProvider(suite.serviceName, suite.jenkinsProvider)
	suite.assert.NotNil(err)
}

func TestCIManager(t *testing.T) {
	suite.Run(t, new(CIManagerTestSuite))
}
