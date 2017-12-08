package goBuilder

import (
	"fmt"

	"github.com/reivaj05/micro-gen/generator/utils"
)

var goFileOptions = []*utils.GenerateFileOptions{
	utils.CreateFileOptions("main.go", "", "main.gen",
		"src/", "go", true),
	utils.CreateFileOptions("endpoints.go", "", "endpoints.gen",
		"src/", "go", false),
	utils.CreateFileOptions("controllers.go", "", "controllers.gen",
		"src/", "go", false),
	utils.CreateFileOptions("config.json", "", "config.gen",
		"src/", "go", false),
}

var buildFileOptions = []*utils.GenerateFileOptions{
	utils.CreateFileOptions("Dockerfile", "", "Dockerfile.gen", "build/",
		"go", true),
	utils.CreateFileOptions("docker-compose.yml", "", "docker-compose.gen", "build/",
		"go", true),
	utils.CreateFileOptions("Makefile", "", "Makefile.gen", "build/",
		"go", false),
	utils.CreateFileOptions("glide.yaml", "", "glide.gen", "build/",
		"go", true),
	utils.CreateFileOptions(".gitignore", "", "ignore.gen", "",
		"go", false),
	utils.CreateFileOptions(".dockerignore", "", "ignore.gen", "",
		"go", false),
	utils.CreateFileOptions(".travis.yml", "", "travis.gen", "build/",
		"go", false),
}

var scriptFileOptions = []*utils.GenerateFileOptions{
	utils.CreateFileOptions("deps.sh", "scripts/", "deps.gen", "scripts/",
		"go", false),
	utils.CreateFileOptions("linter.sh", "scripts/", "linter.gen", "scripts/",
		"go", true),
	utils.CreateFileOptions("start.sh", "scripts/", "start.gen", "scripts/",
		"go", true),
	utils.CreateFileOptions("tests.sh", "scripts/", "tests.gen", "scripts/",
		"go", false),
	utils.CreateFileOptions("build_image.sh", "scripts/", "build_image.gen", "scripts/",
		"go", true),
	utils.CreateFileOptions("push_image.sh", "scripts/", "push_image.gen", "scripts/",
		"go", true),
}

func Build(serviceName string) error {
	if err := createDirectories(serviceName); err != nil {
		return err
	}
	return createService(serviceName)
}

func createDirectories(serviceName string) error {
	paths := []string{fmt.Sprintf("./%s", serviceName), fmt.Sprintf("./%s/scripts", serviceName)}
	for _, path := range paths {
		if err := utils.CreateDir(path); err != nil {
			return err
		}
	}
	return nil
}

func createService(serviceName string) error {
	return generateAllFiles(serviceName)
}

func generateAllFiles(serviceName string) error {
	for _, optionsList := range [][]*utils.GenerateFileOptions{
		goFileOptions, buildFileOptions, scriptFileOptions}{
		if err := generateFilesWithOptionsList(serviceName, optionsList); err != nil {
			return err
		}
	}
	return nil
}

func generateFilesWithOptionsList(serviceName string, fileOptions []*utils.GenerateFileOptions) error {
	for _, options := range fileOptions {
		if err := utils.GenerateFile(serviceName, options); err != nil {
			return err
		}
	}
	return nil
}