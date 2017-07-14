package goBuilder

// TODO: Update/Fix glide package path

import (
	"fmt"

	"github.com/reivaj05/micro-gen/generator/utils"
	"github.com/serenize/snaker"
)

type data struct {
	ServiceName      string
	SnakeServiceName string
}

func Build(serviceName string) error {
	path := fmt.Sprintf("./%s", serviceName)
	if err := utils.CreateDir(path); err != nil {
		return err
	}
	return createService(serviceName)
}

func createService(serviceName string) error {
	return generateFiles(serviceName)
}

func generateFiles(serviceName string) error {
	if err := generateGoFiles(serviceName); err != nil {
		return err
	}
	if err := generateConfigFile(serviceName); err != nil {
		return err
	}
	return generateBuildFiles(serviceName)
}

func generateGoFiles(serviceName string) error {
	if err := generateMainFile(serviceName); err != nil {
		return err
	}
	if err := generateEndpointsFile(serviceName); err != nil {
		return err
	}
	return generateControllersFile(serviceName)
}

func generateMainFile(serviceName string) error {
	return utils.GenerateFile(serviceName, "main", "go",
		"main.gen", "src/", "go", &data{ServiceName: serviceName})
}

func generateEndpointsFile(serviceName string) error {
	return utils.GenerateFile(serviceName, "endpoints", "go",
		"endpoints.gen", "src/", "go", nil)
}

func generateControllersFile(serviceName string) error {
	return utils.GenerateFile(serviceName, "controllers", "go",
		"controllers.gen", "src/", "go", nil)
}

func generateConfigFile(serviceName string) error {
	return utils.GenerateFile(serviceName, "config", "json",
		"config.gen", "src/", "go", nil)
}

func generateBuildFiles(serviceName string) error {
	if err := generateDockerFile(serviceName); err != nil {
		return err
	}
	if err := generateDockerComposeFile(serviceName); err != nil {
		return err
	}
	if err := generateMakeFile(serviceName); err != nil {
		return err
	}
	if err := generateGlideFile(serviceName); err != nil {
		return err
	}
	if err := generateGitIgnoreFile(serviceName); err != nil {
		return err
	}
	if err := generateDockerIgnoreFile(serviceName); err != nil {
		return err
	}
	return generateTravisFile(serviceName)
}

func generateDockerComposeFile(serviceName string) error {
	return utils.GenerateFile(serviceName, "docker-compose", "yml",
		"docker-compose.gen", "build/", "go", &data{
			ServiceName:      serviceName,
			SnakeServiceName: snaker.CamelToSnake(serviceName),
		})
}

func generateDockerFile(serviceName string) error {
	return utils.GenerateFile(serviceName, "Dockerfile", "",
		"Dockerfile.gen", "build/", "go", &data{ServiceName: serviceName})
}

func generateMakeFile(serviceName string) error {
	return utils.GenerateFile(serviceName, "Makefile", "",
		"Makefile.gen", "build/", "go", nil)
}

func generateGlideFile(serviceName string) error {
	return utils.GenerateFile(serviceName, "glide", "yaml",
		"glide.gen", "build/", "go", &data{ServiceName: serviceName})
}

func generateGitIgnoreFile(serviceName string) error {
	return utils.GenerateFile(serviceName, ".gitignore", "",
		"ignore.gen", "", "go", nil)
}

func generateDockerIgnoreFile(serviceName string) error {
	return utils.GenerateFile(serviceName, ".dockerignore", "",
		"ignore.gen", "", "go", nil)
}

func generateTravisFile(serviceName string) error {
	return utils.GenerateFile(serviceName, ".travis", "yml",
		"travis.gen", "build/", "go", nil)
}
