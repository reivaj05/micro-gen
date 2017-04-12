package goGenerator

import (
	"fmt"
	"os"
	"text/template"

	"github.com/reivaj05/micro-gen/generator/utils"
)

type data struct {
	ServiceName string
}

func Build(serviceName string) error {
	fmt.Println("TODO: Implement go generator")
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
	return generateFiles(serviceName)
}

func generateFiles(serviceName string) error {
	if err := generateGoFiles(serviceName); err != nil {
		return err
	}
	if err := generateConfigFile(serviceName); err != nil {
		return err
	}
	return generateBuildFiles(serviceName)
}

func generateGoFiles(serviceName string) error {
	if err := generateMainFile(serviceName); err != nil {
		return err
	}
	if err := generateEndpointsFile(serviceName); err != nil {
		return err
	}
	return generateControllersFile(serviceName)
}

func generateMainFile(serviceName string) error {
	return _generateFile(&utils.GenerateOptions{
		ServiceName:   serviceName,
		FileName:      "main",
		FileExtension: "go",
		FileTemplate:  "main.gen",
		Data: &data{
			ServiceName: serviceName,
		},
	})
}

func generateEndpointsFile(serviceName string) error {
	return _generateFile(&utils.GenerateOptions{
		ServiceName:   serviceName,
		FileName:      "endpoints",
		FileExtension: "go",
		FileTemplate:  "endpoints.gen",
	})
}

func generateControllersFile(serviceName string) error {
	return _generateFile(&utils.GenerateOptions{
		ServiceName:   serviceName,
		FileName:      "controllers",
		FileExtension: "go",
		FileTemplate:  "controllers.gen",
	})
}

func generateConfigFile(serviceName string) error {
	return _generateFile(&utils.GenerateOptions{
		ServiceName:   serviceName,
		FileName:      "config",
		FileExtension: "json",
		FileTemplate:  "config.gen",
	})
}

func generateBuildFiles(serviceName string) error {
	if err := generateDockerFile(serviceName); err != nil {
		return err
	}
	if err := generateMakeFile(serviceName); err != nil {
		return err
	}
	return nil
}

func generateDockerFile(serviceName string) error {
	return _generateFile(&utils.GenerateOptions{
		ServiceName:   serviceName,
		FileName:      "Dockerfile",
		FileExtension: "",
		FileTemplate:  "Dockerfile.gen",
		Data: &data{
			ServiceName: serviceName,
		},
	})
}

func generateMakeFile(serviceName string) error {
	return _generateFile(&utils.GenerateOptions{
		ServiceName:   serviceName,
		FileName:      "Makefile",
		FileExtension: "",
		FileTemplate:  "Makefile.gen",
	})
}

func _generateFile(options *utils.GenerateOptions) error {
	file, err := _createFile(options)
	if err != nil {
		return err
	}
	return _writeTemplateContent(file, options)
}

func _createFile(options *utils.GenerateOptions) (*os.File, error) {
	dst := fmt.Sprintf("./%s/%s", options.ServiceName, options.FileName)
	if options.FileExtension != "" {
		dst = dst + fmt.Sprintf(".%s", options.FileExtension)
	}
	return os.Create(dst)
}

func _writeTemplateContent(file *os.File, options *utils.GenerateOptions) error {
	defer file.Close()
	templateDir := fmt.Sprintf(
		"%s/generator/builders/go/templates/%s",
		utils.GetMicroGenPath(), options.FileTemplate)
	tmpl := template.Must(template.ParseFiles(templateDir))
	return tmpl.Execute(file, options.Data)
}
