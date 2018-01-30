package utils

import (
	"fmt"
	"os"
	"text/template"

	"github.com/reivaj05/GoConfig"

	"github.com/serenize/snaker"
)

// GenerateFileOptions options to create generated files from templates
type GenerateFileOptions struct {
	FileName         string
	FilePath         string
	TemplateFileName string
	TemplateFilePath string
	Language         string
	Data             interface{}
	HasTemplateData  bool
}

type templateData struct {
	ServiceName      string
	SnakeServiceName string
}

// CreateDir creates a new dir in the path passed as parameter
func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// CreateFileOptions creates a new GenerateFileObjects item
func CreateFileOptions(fileName, filePath, templateFileName,
	templateFilePath, language string, hasData bool) *GenerateFileOptions {

	return &GenerateFileOptions{
		FileName:         fileName,
		FilePath:         filePath,
		TemplateFileName: templateFileName,
		TemplateFilePath: templateFilePath,
		Language:         language,
		HasTemplateData:  hasData,
	}
}

// GenerateFile generates a new file from a template
func GenerateFile(serviceName string, options *GenerateFileOptions) error {
	if options.HasTemplateData {
		options.Data = updateTemplateData(serviceName)
	}
	file, err := createFile(serviceName, options)
	if err != nil {
		return err
	}
	return writeTemplateContent(file, options)
}

func updateTemplateData(serviceName string) *templateData {
	return &templateData{
		ServiceName:      serviceName,
		SnakeServiceName: snaker.CamelToSnake(serviceName),
	}
}

func createFile(serviceName string, options *GenerateFileOptions) (*os.File, error) {
	dst := fmt.Sprintf("./%s/%s", serviceName, options.FilePath+options.FileName)
	return os.Create(dst)
}

func writeTemplateContent(file *os.File, options *GenerateFileOptions) error {
	defer file.Close()
	templateDir := getTemplateDir(options)
	if _, err := os.Stat(templateDir); err != nil {
		return err
	}
	tmpl := createTemplate(options.TemplateFileName, templateDir)
	return tmpl.Execute(file, options.Data)
}

func getTemplateDir(options *GenerateFileOptions) string {
	templatesPath := GoConfig.GetConfigMapValue("templates")[options.Language]
	templateFilePath := options.TemplateFilePath
	templateFileName := options.TemplateFileName
	return fmt.Sprintf("%s/%s%s", getMicroGenPath(), templatesPath, templateFilePath+templateFileName)
}

func getMicroGenPath() string {
	const relativePath = "/src/github.com/reivaj05/micro-gen"
	goPath := os.Getenv("GOPATH")
	return goPath + relativePath
}

func createTemplate(filename, templateDir string) *template.Template {
	tmpl := template.New(filename).Funcs(addTemplateFunctions())
	return template.Must(tmpl.ParseFiles(templateDir))
}
func addTemplateFunctions() template.FuncMap {
	return template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}
}
