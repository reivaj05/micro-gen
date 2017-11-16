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
	for _, optionsList := range [][]*utils.GenerateFileOptions{appFileOptions}{
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
