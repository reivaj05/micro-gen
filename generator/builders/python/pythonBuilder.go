package pythonBuilder

import (
	"fmt"

	"github.com/reivaj05/micro-gen/generator/utils"
)

var configFileOptions = []*utils.GenerateFileOptions{
	utils.CreateFileOptions("settings.py", "config/", "settings.gen",
		"src/config/", "python", false),
	utils.CreateFileOptions("urls.py", "config/", "urls.gen", "src/config/",
		"python", false),
	utils.CreateFileOptions("wsgi.py", "config/", "wsgi.gen", "src/config/",
		"python", false),
	utils.CreateFileOptions("__init__.py", "config/", "__init__.gen", "src/",
		"python", false),
	utils.CreateFileOptions("manage.py", "", "manage.gen", "src/",
		"python", false),
}

var appFileOptions = []*utils.GenerateFileOptions{
	utils.CreateFileOptions("apps.py", "app/", "apps.gen", "src/app/",
		"python", false),
	utils.CreateFileOptions("urls.py", "app/", "urls.gen", "src/app/",
		"python", false),
	utils.CreateFileOptions("views.py", "app/", "views.gen", "src/app/",
		"python", false),
	utils.CreateFileOptions("__init__.py", "app/", "__init__.gen", "src/",
		"python", false),
}

var buildFileOptions = []*utils.GenerateFileOptions{
	utils.CreateFileOptions("nose.cfg", "", "nose.gen", "build/",
		"python", false),
	utils.CreateFileOptions("Dockerfile", "", "Dockerfile.gen", "build/",
		"python", true),
	utils.CreateFileOptions("docker-compose.yml", "", "docker-compose.gen",
		"build/", "python", true),
	utils.CreateFileOptions("Makefile", "", "Makefile.gen",
		"build/", "python", true),
	utils.CreateFileOptions("requirements.txt", "", "requirements.gen", "build/",
		"python", false),
	utils.CreateFileOptions(".gitignore", "", "ignore.gen", "",
		"python", false),
	utils.CreateFileOptions(".dockerignore", "", "ignore.gen", "",
		"python", false),
	utils.CreateFileOptions(".travis.yml", "", "travis.gen", "build/",
		"python", false),
}

var scriptFileOptions = []*utils.GenerateFileOptions{
	utils.CreateFileOptions("start.sh", "scripts/", "start.gen", "scripts/",
		"python", false),
	utils.CreateFileOptions("linter.sh", "scripts/", "linter.gen", "scripts/",
		"python", true),
	utils.CreateFileOptions("tests.sh", "scripts/", "tests.gen", "scripts/",
		"python", false),
	utils.CreateFileOptions("deps.sh", "scripts/", "deps.gen", "scripts/",
		"python", false),
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
	if err := generateConfigFiles(serviceName); err != nil {
		return err
	}
	if err := generateAppFiles(serviceName); err != nil {
		return err
	}
	if err := generateBuildFiles(serviceName); err != nil {
		return err
	}
	return generateScriptFiles(serviceName)
}

func generateConfigFiles(serviceName string) error {
	path := fmt.Sprintf("./%s/config", serviceName)
	if err := utils.CreateDir(path); err != nil {
		return err
	}
	for _, options := range configFileOptions {
		if err := utils.GenerateFile(serviceName, options); err != nil {
			return err
		}
	}
	return nil
}

func generateAppFiles(serviceName string) error {
	path := fmt.Sprintf("./%s/app", serviceName)
	if err := utils.CreateDir(path); err != nil {
		return err
	}
	for _, options := range appFileOptions {
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
