package rustGenerator

import (
	"fmt"

	"github.com/reivaj05/micro-gen/generator/utils"
)

var appFileOptions = []*utils.GenerateFileOptions{
	utils.CreateFileOptions("main.rs", "src/", "main.gen", "src/",
		"rust", false),
	utils.CreateFileOptions("Cargo.toml", "", "Cargo.gen", "src/",
		"rust", true),
}

var buildFileOptions = []*utils.GenerateFileOptions{
	utils.CreateFileOptions("Dockerfile", "", "Dockerfile.gen", "build/",
		"rust", true),
	utils.CreateFileOptions("docker-compose.yml", "", "docker-compose.gen", "build/",
		"rust", true),
	utils.CreateFileOptions("Makefile", "", "Makefile.gen", "build/",
		"rust", false),
	utils.CreateFileOptions(".gitignore", "", "ignore.gen", "",
		"rust", false),
	utils.CreateFileOptions(".dockerignore", "", "ignore.gen", "",
		"rust", false),
	utils.CreateFileOptions(".travis.yml", "", "travis.gen", "build/",
		"rust", false),
	utils.CreateFileOptions("README.md", "", "README.gen", "build/",
		"rust", false),
}

var scriptFileOptions = []*utils.GenerateFileOptions{
	utils.CreateFileOptions("start.sh", "scripts/", "start.gen", "scripts/",
		"rust", false),
	utils.CreateFileOptions("linter.sh", "scripts/", "linter.gen", "scripts/",
		"rust", true),
	utils.CreateFileOptions("tests.sh", "scripts/", "tests.gen", "scripts/",
		"rust", false),
	utils.CreateFileOptions("deps.sh", "scripts/", "deps.gen", "scripts/",
		"rust", false),
	utils.CreateFileOptions("build_image.sh", "scripts/", "build_image.gen", "scripts/",
		"rust", true),
	utils.CreateFileOptions("push_image.sh", "scripts/", "push_image.gen", "scripts/",
		"rust", true),
}

func Build(serviceName string) error {
	if err := createDirectories(serviceName); err != nil {
		return err
	}
	return createService(serviceName)
}

func createDirectories(serviceName string) error {
	paths := []string{
		fmt.Sprintf("./%s", serviceName), fmt.Sprintf("./%s/src", serviceName),
		fmt.Sprintf("./%s/scripts", serviceName),
	}
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
		appFileOptions, buildFileOptions, scriptFileOptions} {
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
