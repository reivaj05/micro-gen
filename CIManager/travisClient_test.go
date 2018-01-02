package CIManager

import (
	"fmt"
	"os"
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
	//
}

func TestTravisClient(t *testing.T) {
	suite.Run(t, new(TravisClientTestSuite))
}
