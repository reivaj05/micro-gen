package pythonGenerator

import (
	"fmt"

	"github.com/reivaj05/micro-gen/generator/utils"
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
		return err
	}
	if err := generateURLSAppFile(serviceName); err != nil {
		return err
	}
	if err := generateInitAppFile(serviceName); err != nil {
		return err
	}
	return generateViewsFile(serviceName)
}

func generateAppsFile(serviceName string) error {
	return utils.GenerateFile(serviceName, serviceName+"/app/apps", "py",
		"apps.gen", "src/app/", "python", &data{ServiceName: serviceName})
}

func generateURLSAppFile(serviceName string) error {
	return utils.GenerateFile(serviceName, serviceName+"/app/urls", "py",
		"urls.gen", "src/app/", "python", &data{ServiceName: serviceName})
}

func generateBuildFiles(serviceName string) error {
	// TODO: Implement
	return nil
}
