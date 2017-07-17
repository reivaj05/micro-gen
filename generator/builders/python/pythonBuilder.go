package pythonGenerator

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
	fmt.Println("TODO: Implement python generator")
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
	if err := generateProjectFiles(serviceName); err != nil {
		return err
	}
	if err := generateAppFiles(serviceName); err != nil {
		return err
	}
	return generateBuildFiles(serviceName)
}

func generateProjectFiles(serviceName string) error {
	path := fmt.Sprintf("./%s/%s", serviceName, serviceName)
	if err := utils.CreateDir(path); err != nil {
		return err
	}
	if err := generateSettingsFile(serviceName); err != nil {
		return err
	}
	if err := generateUrlsFile(serviceName); err != nil {
		return err
	}
	if err := generateWSGIFile(serviceName); err != nil {
		return err
	}
	if err := generateInitFile(serviceName); err != nil {
		return err
	}
	return generateManageFile(serviceName)
}

func generateSettingsFile(serviceName string) error {
	return utils.GenerateFile(serviceName, serviceName+"/settings", "py",
		"settings.gen", "src/project/", "python", &data{ServiceName: serviceName})
}

func generateUrlsFile(serviceName string) error {
	return utils.GenerateFile(serviceName, serviceName+"/urls", "py",
		"urls.gen", "src/project/", "python", &data{ServiceName: serviceName})
}

func generateWSGIFile(serviceName string) error {
	return utils.GenerateFile(serviceName, serviceName+"/wsgi", "py",
		"wsgi.gen", "src/project/", "python", &data{ServiceName: serviceName})
}

func generateInitFile(serviceName string) error {
	return utils.GenerateFile(serviceName, serviceName+"/__init__", "py",
		"__init__.gen", "src/", "python", &data{ServiceName: serviceName})
}

func generateManageFile(serviceName string) error {
	return utils.GenerateFile(serviceName, "manage", "py",
		"manage.gen", "src/", "python", &data{ServiceName: serviceName})
}

func generateAppFiles(serviceName string) error {
	path := fmt.Sprintf("./%s/app", serviceName)
	if err := utils.CreateDir(path); err != nil {
		return err
	}
	if err := generateAppsFile(serviceName); err != nil {
		fmt.Println("appsfile")
		return err
	}
	if err := generateURLSAppFile(serviceName); err != nil {
		fmt.Println("urlsfile")
		return err
	}
	if err := generateInitAppFile(serviceName); err != nil {
		fmt.Println("initfile")
		return err
	}
	return generateViewsFile(serviceName)
}

func generateAppsFile(serviceName string) error {
	return utils.GenerateFile(serviceName, "/app/apps", "py",
		"apps.gen", "src/app/", "python", &data{ServiceName: serviceName})
}

func generateURLSAppFile(serviceName string) error {
	return utils.GenerateFile(serviceName, "/app/urls", "py",
		"urls.gen", "src/app/", "python", &data{ServiceName: serviceName})
}

func generateInitAppFile(serviceName string) error {
	return utils.GenerateFile(serviceName, "/app/__init__", "py",
		"__init__.gen", "src/", "python", &data{ServiceName: serviceName})
}

func generateViewsFile(serviceName string) error {
	return utils.GenerateFile(serviceName, "/app/views", "py",
		"views.gen", "src/app/", "python", &data{ServiceName: serviceName})
}

func generateBuildFiles(serviceName string) error {
	if err := generateDockerFile(serviceName); err != nil {
		return err
	}
	if err := generateDockerComposeFile(serviceName); err != nil {
		return err
	}
	if err := generateRequirementsFile(serviceName); err != nil {
		return err
	}
	if err := generateGitIgnoreFile(serviceName); err != nil {
		return err
	}
	if err := generateDockerIgnoreFile(serviceName); err != nil {
		return err
	}
	return nil
	// return generateTravisFile(serviceName)
}

func generateDockerFile(serviceName string) error {
	return utils.GenerateFile(serviceName, "Dockerfile", "",
		"Dockerfile.gen", "build/", "python", &data{ServiceName: serviceName})
}

func generateDockerComposeFile(serviceName string) error {
	return utils.GenerateFile(serviceName, "docker-compose", "yml",
		"docker-compose.gen", "build/", "python", &data{
			ServiceName:      serviceName,
			SnakeServiceName: snaker.CamelToSnake(serviceName),
		})
}

func generateRequirementsFile(serviceName string) error {
	return utils.GenerateFile(serviceName, "requirements", "txt",
		"requirements.gen", "build/", "python", nil)
}

func generateGitIgnoreFile(serviceName string) error {
	return utils.GenerateFile(serviceName, ".gitignore", "",
		"ignore.gen", "", "python", &data{ServiceName: serviceName})
}

func generateDockerIgnoreFile(serviceName string) error {
	return utils.GenerateFile(serviceName, ".dockerignore", "",
		"ignore.gen", "", "python", nil)
}
