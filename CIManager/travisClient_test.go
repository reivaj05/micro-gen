package CIManager

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TravisClientTestSuite struct {
	suite.Suite
	assert      *assert.Assertions
	serviceName string
	token       string
	mockServer  *httptest.Server
}

type mockHandler struct{}

func (suite *TravisClientTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.token = "mockToken"
	suite.serviceName = "mockServiceName"
	suite.mockServer = httptest.NewServer(&mockHandler{})
	baseURL = suite.mockServer.URL
	reposEndpoint = suite.mockServer.URL
	repoActivateEndpoint = suite.mockServer.URL
	userEndpoint = suite.mockServer.URL
	syncAccountEndpoint = suite.mockServer.URL
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

func (suite *mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	SendResponseWithStatus(w, "{}", http.StatusOK)
}

func SendResponseWithStatus(
	w http.ResponseWriter, response string, status int) error {

	w.WriteHeader(status)
	_, err := fmt.Fprintf(w, response)
	return err
}

func TestTravisClient(t *testing.T) {
	suite.Run(t, new(TravisClientTestSuite))
}
