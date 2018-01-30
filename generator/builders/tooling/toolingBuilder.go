package toolingBuilder

import (
	"fmt"
	"os"
	"strings"

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
	return generateAllFiles(serviceName, services)
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
