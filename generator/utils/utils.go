package utils

import (
	"fmt"
	"os"
	"text/template"

	"github.com/reivaj05/GoConfig"
)

type generateOptions struct {
	ServiceName   string
	FileExtension string
	FileTemplate  string
	FileName      string
	FilePath      string
	Data          interface{}
}

func GenerateFile(serviceName, fileName, fileExtension,
	fileTemplate, filePath string, data interface{}) error {

	options := createGenerateOptions(serviceName, fileName,
		fileExtension, fileTemplate, filePath, data)
	file, err := createFile(options)
	if err != nil {
		return err
	}
	return writeTemplateContent(file, options)
}

func createGenerateOptions(serviceName, fileName, fileExtension,
	fileTemplate, filePath string, data interface{}) *generateOptions {

	return &generateOptions{
		ServiceName:   serviceName,
		FileName:      fileName,
		FileExtension: fileExtension,
		FileTemplate:  fileTemplate,
		FilePath:      filePath,
		Data:          data,
	}
}

func createFile(options *generateOptions) (*os.File, error) {
	dst := fmt.Sprintf("./%s/%s", options.ServiceName, options.FileName)
	if options.FileExtension != "" {
		dst = dst + fmt.Sprintf(".%s", options.FileExtension)
	}
	return os.Create(dst)
}

func writeTemplateContent(file *os.File, options *generateOptions) error {
	defer file.Close()
	templateDir := fmt.Sprintf("%s/%s%s%s", getMicroGenPath(),
		GoConfig.GetConfigStringValue("goTemplatesPath"), options.FilePath,
		options.FileTemplate)
	if _, err := os.Stat(templateDir); err != nil {
		return err
	}
	tmpl := template.Must(template.ParseFiles(templateDir))
	return tmpl.Execute(file, options.Data)
}

func getMicroGenPath() string {
	const relativePath = "/src/github.com/reivaj05/micro-gen"
	goPath := os.Getenv("GOPATH")
	return goPath + relativePath
}
