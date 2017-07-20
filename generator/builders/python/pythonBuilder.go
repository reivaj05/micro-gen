package pythonBuilder

import (
	"fmt"

	"github.com/reivaj05/micro-gen/generator/utils"
)

var projectFileOptions = []*utils.GenerateFileOptions{
	utils.CreateFileOptions("settings.py", "config/", "settings.gen", "src/project/",
		"python", false),
	utils.CreateFileOptions("urls.py", "config/", "urls.gen", "src/project/",
		"python", false),
	utils.CreateFileOptions("wsgi.py", "config/", "wsgi.gen", "src/project/",
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
	if err := generateProjectFiles(serviceName); err != nil {
		return err
	}
	if err := generateAppFiles(serviceName); err != nil {
		return err
	}
	return nil
	// if err := generateBuildFiles(serviceName); err != nil {
	// 	return err
	// }
	// return generateScriptFiles(serviceName)
}

func generateProjectFiles(serviceName string) error {
	path := fmt.Sprintf("./%s/config", serviceName)
	if err := utils.CreateDir(path); err != nil {
		return err
	}
	for _, options := range projectFileOptions {
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

// func generateBuildFiles(serviceName string) error {
// 	if err := generateNoseFile(serviceName); err != nil {
// 		return err
// 	}
// 	if err := generateDockerFile(serviceName); err != nil {
// 		return err
// 	}
// 	if err := generateDockerComposeFile(serviceName); err != nil {
// 		return err
// 	}
// 	if err := generateRequirementsFile(serviceName); err != nil {
// 		return err
// 	}
// 	if err := generateGitIgnoreFile(serviceName); err != nil {
// 		return err
// 	}
// 	if err := generateDockerIgnoreFile(serviceName); err != nil {
// 		return err
// 	}
// 	return generateTravisFile(serviceName)
// }

// func generateNoseFile(serviceName string) error {
// 	return utils.GenerateFile(serviceName, "nose", "cfg",
// 		"nose.gen", "build/", "python", &data{ServiceName: serviceName})
// }

// func generateDockerFile(serviceName string) error {
// 	return utils.GenerateFile(serviceName, "Dockerfile", "",
// 		"Dockerfile.gen", "build/", "python", &data{ServiceName: serviceName})
// }

// func generateDockerComposeFile(serviceName string) error {
// 	return utils.GenerateFile(serviceName, "docker-compose", "yml",
// 		"docker-compose.gen", "build/", "python", &data{
// 			ServiceName:      serviceName,
// 			SnakeServiceName: snaker.CamelToSnake(serviceName),
// 		})
// }

// func generateRequirementsFile(serviceName string) error {
// 	return utils.GenerateFile(serviceName, "requirements", "txt",
// 		"requirements.gen", "build/", "python", nil)
// }

// func generateGitIgnoreFile(serviceName string) error {
// 	return utils.GenerateFile(serviceName, ".gitignore", "",
// 		"ignore.gen", "", "python", nil)
// }

// func generateDockerIgnoreFile(serviceName string) error {
// 	return utils.GenerateFile(serviceName, ".dockerignore", "",
// 		"ignore.gen", "", "python", nil)
// }

// func generateTravisFile(serviceName string) error {
// 	return utils.GenerateFile(serviceName, ".travis", "yml",
// 		"travis.gen", "build/", "python", nil)
// }

// func generateScriptFiles(serviceName string) error {
// 	path := fmt.Sprintf("./%s/scripts", serviceName)
// 	if err := utils.CreateDir(path); err != nil {
// 		return err
// 	}
// 	if err := generateStartFile(serviceName); err != nil {
// 		return err
// 	}
// 	if err := generateLinterFile(serviceName); err != nil {
// 		return err
// 	}
// 	return generateTestsFile(serviceName)
// }

// func generateStartFile(serviceName string) error {
// 	return utils.GenerateFile(serviceName, "scripts/start", "sh",
// 		"start.gen", "scripts/", "python", nil)
// }

// func generateLinterFile(serviceName string) error {
// 	return utils.GenerateFile(serviceName, "scripts/linter", "sh",
// 		"linter.gen", "scripts/", "python", &data{ServiceName: serviceName})
// }

// func generateTestsFile(serviceName string) error {
// 	return utils.GenerateFile(serviceName, "scripts/tests", "sh",
// 		"tests.gen", "scripts/", "python", nil)
// }
