package repoManager

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BitbucketClientTestSuite struct {
	suite.Suite
	assert      *assert.Assertions
	serviceName string
	token       string
}

func (suite *BitbucketClientTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.serviceName = "mockService"
	suite.token = "mockToken"
}

func (suite *BitbucketClientTestSuite) TestCreateBitbucketRepoSuccessfully() {
	// client := NewBitbucketClient(suite.token)
	// url, err := client.CreateCloudRepo(suite.serviceName)
	// suite.assert.Nil(err)
	// suite.assert.NotEqual(url, "")
}

func (suite *BitbucketClientTestSuite) TestCreateBitbucketRepoUnsuccessfully() {
	// client := NewBitbucketClient(suite.token)
	// url, err := client.CreateCloudRepo(suite.serviceName)
	// suite.assert.NotNil(err)
	// suite.assert.Equal(url, "")
}

func TestBitbucketClient(t *testing.T) {
	suite.Run(t, new(BitbucketClientTestSuite))
}
