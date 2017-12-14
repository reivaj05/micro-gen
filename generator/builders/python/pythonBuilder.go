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
	utils.CreateFileOptions("build_image.sh", "scripts/", "build_image.gen", "scripts/",
		"python", true),
	utils.CreateFileOptions("push_image.sh", "scripts/", "push_image.gen", "scripts/",
		"python", true),
}

func Build(serviceName string) error {
	if err := createDirectories(serviceName); err != nil {
		return err
	}
	return createService(serviceName)
}

func createDirectories(serviceName string) error {
	paths := []string{
		fmt.Sprintf("./%s", serviceName), fmt.Sprintf("./%s/config", serviceName),
		fmt.Sprintf("./%s/app", serviceName), fmt.Sprintf("./%s/scripts", serviceName),
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
		configFileOptions, appFileOptions, buildFileOptions, scriptFileOptions} {
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
