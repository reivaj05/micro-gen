package pythonGenerator

import (
	"fmt"
	"os"
)

func Build(serviceName string) error {
	fmt.Println("TODO: Implement python generator")
	if err := createServiceDir(serviceName); err != nil {
		return err
	}
	return createService(serviceName)
}

// TODO: Refactor to one function(all builders create a new directory)
func createServiceDir(serviceName string) error {
	dst := fmt.Sprintf("./%s", serviceName)
	return os.MkdirAll(dst, os.ModePerm)
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
