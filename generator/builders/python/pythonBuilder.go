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

func createServiceDir(serviceName string) error {
	dst := fmt.Sprintf("./%s", serviceName)
	return os.MkdirAll(dst, os.ModePerm)
}

func createService(serviceName string) error {
	fmt.Println("TODO: Implement create service function")
	return nil
}
