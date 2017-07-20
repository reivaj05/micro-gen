package goBuilder

// TODO: Update/Fix glide package path
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
	utils.CreateFileOptions("travis.yml", "", "travis.gen", "build/",
		"go", false),
}

var scriptFileOptions = []*utils.GenerateFileOptions{
	// utils.CreateFileOptions("start.sh", "scripts/", "start.gen", "scripts/",
	// 	"go", false),
	// utils.CreateFileOptions("linter.sh", "scripts/", "linter.gen", "scripts/",
	// 	"go", true),
	utils.CreateFileOptions("tests.sh", "scripts/", "tests.gen", "scripts/",
		"go", false),
	utils.CreateFileOptions("deps.sh", "scripts/", "deps.gen", "scripts/",
		"go", false),
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
	if err := generateBuildFiles(serviceName); err != nil {
		return err
	}
	return generateScriptFiles(serviceName)
}

func generateGoFiles(serviceName string) error {
	for _, options := range goFileOptions {
		if err := utils.GenerateFile(serviceName, options); err != nil {
			return err
		}
	}
	return nil
}

func generateBuildFiles(serviceName string) error {
	for _, options := range buildFileOptions {
		if err := utils.GenerateFile(serviceName, options); err != nil {
			return err
		}
	}
	return nil
}

func generateScriptFiles(serviceName string) error {
	path := fmt.Sprintf("./%s/scripts", serviceName)
	if err := utils.CreateDir(path); err != nil {
		return err
	}
	for _, options := range scriptFileOptions {
		if err := utils.GenerateFile(serviceName, options); err != nil {
			return err
		}
	}
	return nil
}
