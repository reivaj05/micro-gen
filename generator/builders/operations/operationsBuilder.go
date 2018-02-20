package operationsBuilder

import (
	"fmt"
	"os"
	"strings"

	"github.com/reivaj05/micro-gen/generator/utils"

	"github.com/reivaj05/micro-gen/docker-wrapper"
)

func Build(opName, services string) error {
	if err := createDirectories(opName); err != nil {
		return err
	}
	return createService(opName, strings.Split(services, ","))
}

func createDirectories(opName string) error {
	paths := []string{fmt.Sprintf("./%s", opName)}
	for _, path := range paths {
		if err := utils.CreateDir(path); err != nil {
			return err
		}
	}
	return nil
}

func createService(opName string, services []string) error {
	services = filterServices(services)
	return generateAllFiles(opName, services)
}

func filterServices(services []string) []string {
	if docker, err := dockerWrapper.NewDockerRegistryManager(); err == nil {
		return docker.FilterByExistingRepos(services)
	}
	return services
}

func generateAllFiles(opName string, services []string) error {
	for index, service := range services {
		if err := generateFilesForService(opName, service, index); err != nil {
			return err
		}
	}
	return nil
}

func generateFilesForService(opName, service string, index int) error {
	if err := utils.CreateDir(fmt.Sprintf("./%s/%s", opName, service)); err != nil {
		return err
	}
	if err := generateDeploymentFile(opName, service); err != nil {
		return err
	}
	return generateServiceFile(opName, service, index)
}

func generateDeploymentFile(opName, service string) error {
	options := utils.CreateFileOptions("deployment.yml", fmt.Sprintf("%s/", service),
		"deployment.gen", "", "operations", false)
	options.Data = createOptionsDeploymentData(service)
	return utils.GenerateFile(opName, options)
}

func createOptionsDeploymentData(service string) interface{} {
	return struct {
		Service        string
		DockerUsername string
	}{
		Service:        service,
		DockerUsername: os.Getenv("DOCKER_USERNAME"),
	}
}

func generateServiceFile(opName, service string, index int) error {
	options := utils.CreateFileOptions("service.yml", fmt.Sprintf("%s/", service),
		"service.gen", "", "operations", false)
	options.Data = createOptionsServicetData(service, index)
	return utils.GenerateFile(opName, options)
}

func createOptionsServicetData(service string, index int) interface{} {
	return struct {
		Service string
		Index   int
	}{
		Service: service,
		Index:   index,
	}
}
