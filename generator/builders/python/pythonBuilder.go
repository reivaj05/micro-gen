package pythonGenerator

import (
	"fmt"

	"github.com/reivaj05/micro-gen/generator/utils"
)

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
	// TODO: Implement
	return nil
}

func generateAppFiles(serviceName string) error {
	// TODO: Implement
	return nil
}

func generateBuildFiles(serviceName string) error {
	// TODO: Implement
	return nil
}
