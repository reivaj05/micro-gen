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
	assert                 *assert.Assertions
	serviceName            string
	token                  string
	mockReposServer        *httptest.Server
	mockRepoActivateServer *httptest.Server
	mockUserServer         *httptest.Server
	mockSyncAccountServer  *httptest.Server
}

type mockReposHandler struct{}
type mockRepoActivateHandler struct{}
type mockUserHandler struct{}
type mockSyncAccountHandler struct{}

type statusObj struct {
	repos, repoActivate, user, syncAccount, notFoundRepo int
}

const successStatus = 1
const failureStatus = 0

var currentStatus *statusObj

func (suite *TravisClientTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.token = "mockToken"
	suite.serviceName = "mockServiceName"
	suite.createMockServers()
	suite.initMockEndpoints()
}

func (suite *TravisClientTestSuite) createMockServers() {
	suite.mockReposServer = httptest.NewServer(&mockReposHandler{})
	suite.mockRepoActivateServer = httptest.NewServer(&mockRepoActivateHandler{})
	suite.mockUserServer = httptest.NewServer(&mockUserHandler{})
	suite.mockSyncAccountServer = httptest.NewServer(&mockSyncAccountHandler{})
}

func (handler *mockReposHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := `{"repositories":[{"name": "mockServiceName", "slug": "mockSlug"}]}`
	if currentStatus.notFoundRepo == failureStatus {
		response = `{"repositories":[{"name": "wrongServiceName"}]}`
	}
	sendResponseWithStatus(w, response, currentStatus.repos)
}

func (handler *mockRepoActivateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sendResponseWithStatus(w, "{}", currentStatus.repoActivate)
}

func (handler *mockUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sendResponseWithStatus(w, `{"id": 1}`, currentStatus.user)
}

func (handler *mockSyncAccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sendResponseWithStatus(w, "{}", currentStatus.syncAccount)
}

func sendResponseWithStatus(w http.ResponseWriter, goodResponse string, status int) {
	if status == failureStatus {
		sendResponse(w, `{}`, http.StatusBadRequest)
		return
	}
	sendResponse(w, goodResponse, http.StatusOK)
}

func sendResponse(w http.ResponseWriter, response string, status int) {
	w.WriteHeader(status)
	fmt.Fprintf(w, response)
}

func (suite *TravisClientTestSuite) initMockEndpoints() {
	baseURL = ""
	reposEndpoint = suite.mockReposServer.URL
	repoActivateEndpoint = fmt.Sprintf("%s?%s%s", suite.mockRepoActivateServer.URL)
	userEndpoint = suite.mockUserServer.URL
	syncAccountEndpoint = fmt.Sprintf("%s?%s%s", suite.mockSyncAccountServer.URL)
}

func (suite *TravisClientTestSuite) TearDownSuite() {
	suite.mockReposServer.Close()
	suite.mockRepoActivateServer.Close()
	suite.mockUserServer.Close()
	suite.mockSyncAccountServer.Close()
}

func (suite *TravisClientTestSuite) TestNewTravisClient() {
	client := NewTravisClient(suite.token)
	suite.assert.NotNil(client)
}

func (suite *TravisClientTestSuite) TestActivateRepoSuccessfully() {
	ss := successStatus
	currentStatus = suite.updateCurrentStatus(ss, ss, ss, ss, ss)
	client := NewTravisClient(suite.token)
	err := client.ActivateRepo(suite.serviceName)
	suite.assert.Nil(err)
}

func (suite *TravisClientTestSuite) TestActivateRepoUserEndpointError() {
	ss := successStatus
	fs := failureStatus
	currentStatus = suite.updateCurrentStatus(ss, ss, fs, ss, ss)
	client := NewTravisClient(suite.token)
	err := client.ActivateRepo(suite.serviceName)
	suite.assert.NotNil(err)
}

func (suite *TravisClientTestSuite) TestActivateRepoWrongCredentials() {

}

func (suite *TravisClientTestSuite) TestActivateRepoSyncAccountEndpointError() {
	ss := successStatus
	fs := failureStatus
	currentStatus = suite.updateCurrentStatus(ss, ss, ss, fs, ss)
	client := NewTravisClient(suite.token)
	err := client.ActivateRepo(suite.serviceName)
	suite.assert.NotNil(err)
}

func (suite *TravisClientTestSuite) TestActivateRepoReposEndpointError() {
	ss := successStatus
	fs := failureStatus
	currentStatus = suite.updateCurrentStatus(fs, ss, ss, ss, ss)
	client := NewTravisClient(suite.token)
	err := client.ActivateRepo(suite.serviceName)
	suite.assert.NotNil(err)
}

func (suite *TravisClientTestSuite) TestActivateRepoRepoNotFoundError() {
	ss := successStatus
	fs := failureStatus
	currentStatus = suite.updateCurrentStatus(ss, ss, ss, ss, fs)
	client := NewTravisClient(suite.token)
	err := client.ActivateRepo(suite.serviceName)
	suite.assert.NotNil(err)
}

func (suite *TravisClientTestSuite) updateCurrentStatus(
	repos, repoActivate, user, syncAccount, notFoundRepo int) *statusObj {

	return &statusObj{
		repos:        repos,
		repoActivate: repoActivate,
		user:         user,
		syncAccount:  syncAccount,
		notFoundRepo: notFoundRepo,
	}
}

func TestTravisClient(t *testing.T) {
	suite.Run(t, new(TravisClientTestSuite))
}
