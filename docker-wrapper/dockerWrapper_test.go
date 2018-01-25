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
}

type loginHandler struct{}

func (suite *DockerWrapperTestSuite) SetupSuite() {
	suite.assert = assert.New(suite.T())
	suite.createMockServers()
}

func (suite *DockerWrapperTestSuite) createMockServers() {
	suite.loginServer = httptest.NewServer(&loginHandler{})
}

func (handler *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := `{"repositories":[{"name": "mockServiceName", "slug": "mockSlug"}]}`
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, response)
}

func (suite *DockerWrapperTestSuite) TearDownSuite() {
	suite.loginServer.Close()
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
	loginEndpoint = fmt.Sprintf("%s?%s", suite.loginServer.URL)
	suite.setupEnvVars("MOCK_USERNAME", "MOCK_PASSWORD", "MOCK_REGISTRY")
	manager, err := NewDockerRegistryManager()
	suite.assert.NotNil(manager)
	suite.assert.Nil(err)
}

func (suite *DockerWrapperTestSuite) TestNewDockerRegistryManagerWithoutDockerUsername() {
	suite.setupEnvVars("", "", "")
	manager, err := NewDockerRegistryManager()
	suite.assert.Nil(manager)
	suite.assert.NotNil(err)

}
func (suite *DockerWrapperTestSuite) TestNewDockerRegistryManagerWithoutDockerPasword() {
	suite.setupEnvVars("MOCK_USERNAME", "", "")
	manager, err := NewDockerRegistryManager()
	suite.assert.Nil(manager)
	suite.assert.NotNil(err)

}
func (suite *DockerWrapperTestSuite) TestNewDockerRegistryManagerWithoutDockerRegistryHost() {
	suite.setupEnvVars("MOCK_USERNAME", "MOCK_PASSWORD", "")
	manager, err := NewDockerRegistryManager()
	suite.assert.Nil(manager)
	suite.assert.NotNil(err)

}

func (suite *DockerWrapperTestSuite) TestNewDockerRegistryManagerErrorInLogin() {
	// TODO: Add failing server for coverage
	suite.setupEnvVars("MOCK_USERNAME", "MOCK_PASSWORD", "MOCK_REGISTRY")
	manager, err := NewDockerRegistryManager()
	suite.assert.Nil(manager)
	suite.assert.NotNil(err)
}

func (suite *DockerWrapperTestSuite) setupEnvVars(username, password, registry string) {
	os.Setenv(dockerUsernameKey, username)
	os.Setenv(dockerPasswordKey, password)
	os.Setenv(dockerRegistryHostKey, registry)
}

func TestDockerWrapper(t *testing.T) {
	suite.Run(t, new(DockerWrapperTestSuite))
}
