package CIManager

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TravisClientTestSuite struct {
	suite.Suite
	assert      *assert.Assertions
	serviceName string
	token       string
}

func (suite *TravisClientTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.token = "mockToken"
	suite.serviceName = "mockServiceName"
}

func (suite *TravisClientTestSuite) SetupTest() {
	//
}

func (suite *TravisClientTestSuite) TearDownTest() {
	//
}

func (suite *TravisClientTestSuite) TestNewTravisClient() {
	client := NewTravisClient(suite.token)
	suite.assert.NotNil(client)
}

func (suite *TravisClientTestSuite) TestActivateRepoSuccessfully() {

}

func (suite *TravisClientTestSuite) TestActivateRepoUserEndpointError() {

}

func (suite *TravisClientTestSuite) TestActivateRepoWrongCredentials() {

}

func (suite *TravisClientTestSuite) TestActivateRepoSyncAccountEndpointError() {

}

func (suite *TravisClientTestSuite) TestActivateRepoReposEndpointError() {

}

func (suite *TravisClientTestSuite) TestActivateRepoRepoNotFoundError() {

}

func TestTravisClient(t *testing.T) {
	suite.Run(t, new(TravisClientTestSuite))
}
