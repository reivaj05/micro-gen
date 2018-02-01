package toolingBuilder

import (
	"fmt"
	"github.com/reivaj05/GoJSON"
	"os"
	"strings"

	"github.com/reivaj05/micro-gen/docker-wrapper"
	"github.com/reivaj05/micro-gen/generator/utils"
)

var buildFileOptions = []*utils.GenerateFileOptions{
	utils.CreateFileOptions("docker-compose.yml", "", "docker-compose.gen", "build/",
		"tooling", false),
}

func Build(serviceName, services string) error {
	if err := createDirectories(serviceName); err != nil {
		return err
	}
	return createService(serviceName, services)
}

func createDirectories(serviceName string) error {
	paths := []string{fmt.Sprintf("./%s", serviceName)}
	for _, path := range paths {
		if err := utils.CreateDir(path); err != nil {
			return err
		}
	}
	return nil
}

func createService(serviceName, services string) error {
	services = filterServices(strings.Split(services, ","))
	return generateAllFiles(serviceName, services)
}

func filterServices(services []string) string {
	if docker, err := dockerWrapper.NewDockerRegistryManager(); err == nil {
		reposResponse, err := docker.SearchRepos()
		if err == nil {
			services = filterAgainstDockerRegistryRepos(services, reposResponse)
		}
	}
	return strings.Join(services, ",")
}

func filterAgainstDockerRegistryRepos(
	services []string, reposResponse *GoJSON.JSONWrapper) (filteredServices []string) {

	repos := reposResponse.GetArrayFromPath("results")
	for _, service := range services {
		if serviceIsInDockerRegistry(repos, service) {
			filteredServices = append(filteredServices, service)
		}
	}
	return filteredServices
}

func generateAllFiles(serviceName, services string) error {
	for _, optionsList := range [][]*utils.GenerateFileOptions{buildFileOptions} {
		if err := generateFilesWithOptionsList(serviceName, services, optionsList); err != nil {
			return err
		}
	}
	return nil
}

func generateFilesWithOptionsList(serviceName, services string, fileOptions []*utils.GenerateFileOptions) error {
	for _, options := range fileOptions {
		options.Data = struct {
			DockerUsername string
			Services       []string
		}{
			DockerUsername: os.Getenv("DOCKER_USERNAME"),
			Services:       strings.Split(services, ","),
		}
		if err := utils.GenerateFile(serviceName, options); err != nil {
			return err
		}
	}
	return nil
}
