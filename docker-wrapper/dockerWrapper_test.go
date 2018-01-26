package dockerWrapper

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DockerWrapperTestSuite struct {
	suite.Suite
	assert             *assert.Assertions
	dockerUsername     string
	dockerPassword     string
	dockerRegistryHost string
	loginServer        *httptest.Server
	reposServer        *httptest.Server
}

type loginHandler struct{}
type reposHandler struct{}

var successStatus = 0
var failureStatus = 1
var currentStatus int

func (suite *DockerWrapperTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.createMockServers()
}

func (suite *DockerWrapperTestSuite) createMockServers() {
	suite.loginServer = httptest.NewServer(&loginHandler{})
	suite.reposServer = httptest.NewServer(&reposHandler{})
}

func (handler *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, `{"token": "mockToken"}`)
}

func (handler *reposHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, `{"results":[]}`)
}

func sendResponse(w http.ResponseWriter, response string) {
	var status int
	if currentStatus == successStatus {
		status = http.StatusOK
	} else {
		status = http.StatusBadRequest
	}
	w.WriteHeader(status)
	fmt.Fprintf(w, response)
}

func (suite *DockerWrapperTestSuite) TearDownSuite() {
	suite.loginServer.Close()
	suite.reposServer.Close()
}

func (suite *DockerWrapperTestSuite) SetupTest() {
	suite.dockerUsername = os.Getenv(dockerUsernameKey)
	suite.dockerPassword = os.Getenv(dockerPasswordKey)
	suite.dockerRegistryHost = os.Getenv(dockerRegistryHostKey)
}

func (suite *DockerWrapperTestSuite) TearDownTest() {
	os.Setenv(dockerUsernameKey, suite.dockerUsername)
	os.Setenv(dockerPasswordKey, suite.dockerPassword)
	os.Setenv(dockerRegistryHostKey, suite.dockerRegistryHost)
}

func (suite *DockerWrapperTestSuite) TestNewDockerRegistryManagerSuccessful() {
	currentStatus = successStatus
	loginEndpoint = fmt.Sprintf("%s?%s", suite.loginServer.URL)
	suite.setupEnvVars("MOCK_USERNAME", "MOCK_PASSWORD", "MOCK_REGISTRY")
	manager, err := NewDockerRegistryManager()
	suite.assert.NotNil(manager)
	suite.assert.Nil(err)
}

func (suite *DockerWrapperTestSuite) TestNewDockerRegistryManagerWithoutDockerUsername() {
	suite.setupEnvVars("", "", "")
	suite.assertWrongNewDockerRegistryManager()

}
func (suite *DockerWrapperTestSuite) TestNewDockerRegistryManagerWithoutDockerPasword() {
	suite.setupEnvVars("MOCK_USERNAME", "", "")
	suite.assertWrongNewDockerRegistryManager()

}
func (suite *DockerWrapperTestSuite) TestNewDockerRegistryManagerWithoutDockerRegistryHost() {
	suite.setupEnvVars("MOCK_USERNAME", "MOCK_PASSWORD", "")
	suite.assertWrongNewDockerRegistryManager()
}

func (suite *DockerWrapperTestSuite) TestNewDockerRegistryManagerErrorInLogin() {
	currentStatus = failureStatus
	suite.setupEnvVars("MOCK_USERNAME", "MOCK_PASSWORD", "MOCK_REGISTRY")
	suite.assertWrongNewDockerRegistryManager()
}

func (suite *DockerWrapperTestSuite) assertWrongNewDockerRegistryManager() {
	manager, err := NewDockerRegistryManager()
	suite.assert.Nil(manager)
	suite.assert.NotNil(err)
}

func (suite *DockerWrapperTestSuite) TestSearchReposSuccessful() {
	currentStatus = successStatus
	repositoriesEndpoint = fmt.Sprintf("%s?%s%s", suite.reposServer.URL)
	suite.setupEnvVars("MOCK_USERNAME", "MOCK_PASSWORD", "MOCK_REGISTRY")
	manager, _ := NewDockerRegistryManager()
	data, err := manager.SearchRepos()
	suite.assert.NotNil(data)
	suite.assert.Nil(err)
	suite.assert.True(data.HasPath("results"))
}

func (suite *DockerWrapperTestSuite) TestSearchReposUnsuccessful() {
	currentStatus = failureStatus
	repositoriesEndpoint = fmt.Sprintf("%s?%s%s", suite.reposServer.URL)
	suite.setupEnvVars("MOCK_USERNAME", "MOCK_PASSWORD", "MOCK_REGISTRY")
	manager, _ := NewDockerRegistryManager()
	data, err := manager.SearchRepos()
	suite.assert.Nil(data)
	suite.assert.NotNil(err)
}

func (suite *DockerWrapperTestSuite) setupSearchRepos(status int) {
	currentStatus = status
	repositoriesEndpoint = fmt.Sprintf("%s?%s%s", suite.reposServer.URL)
	suite.setupEnvVars("MOCK_USERNAME", "MOCK_PASSWORD", "MOCK_REGISTRY")
}

func (suite *DockerWrapperTestSuite) setupEnvVars(username, password, registry string) {
	os.Setenv(dockerUsernameKey, username)
	os.Setenv(dockerPasswordKey, password)
	os.Setenv(dockerRegistryHostKey, registry)
}

func TestDockerWrapper(t *testing.T) {
	suite.Run(t, new(DockerWrapperTestSuite))
}
