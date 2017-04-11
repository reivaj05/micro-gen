package goGenerator

import "fmt"

func Build(serviceName string) error {
	fmt.Println("TODO: Implement go generator")
	return nil
}

// func createService(args ...string) {
// 	// TODO: Create properly template files (protos and go files for services)
// 	// TODO: Implement rest of template files
// 	basePath := joinPath()
// 	for _, serviceName := range args {
// 		if err := generateFiles(basePath, serviceName); err != nil {
// 			fmt.Errorf("Service not created: " + err.Error())
// 			rollback(serviceName)
// 		}
// 	}

// }

// func joinPath() string {
// 	const relativePath = "/src/github.com/reivaj05/apigateway"
// 	goPath := os.Getenv("GOPATH")
// 	return goPath + relativePath
// }

// func generateFiles(path, serviceName string) error {
// 	if err := generateAPIFile(path, serviceName); err != nil {
// 		return err
// 	}
// 	if err := generateServiceFile(path, serviceName); err != nil {
// 		return err
// 	}
// 	if err := generateProtoFiles(path, serviceName); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func generateAPIFile(path, serviceName string) error {
// 	path += "/api/" + serviceName + "/"
// 	return _generateFile(&generateOptions{
// 		path:          path,
// 		serviceName:   serviceName,
// 		fileExtension: ".go",
// 		fileTemplate:  "goAPI.txt",
// 		data: &goAPITemplateData{
// 			ServiceName:      serviceName,
// 			UpperServiceName: inflect.Titleize(serviceName),
// 		},
// 	})

// }

// func generateServiceFile(path, serviceName string) error {
// 	path += "/services/" + serviceName + "/"
// 	return _generateFile(&generateOptions{
// 		path:          path,
// 		serviceName:   serviceName,
// 		fileExtension: ".go",
// 		fileTemplate:  "goService.txt",
// 		data: &goAPITemplateData{
// 			ServiceName:      serviceName,
// 			UpperServiceName: inflect.Titleize(serviceName),
// 		},
// 	})
// }

// func generateProtoFiles(path, serviceName string) error {
// 	sp := camelcase.Split(serviceName)[0]
// 	if err := _generateFile(&generateOptions{
// 		path:          path + "/protos/api/",
// 		serviceName:   serviceName,
// 		fileExtension: ".proto",
// 		fileTemplate:  "protoAPI.txt",
// 		data: &protoAPITemplateData{
// 			ServiceName:      serviceName,
// 			ResourcePath:     sp,
// 			UpperServiceName: inflect.Titleize(serviceName),
// 		},
// 	}); err != nil {
// 		return err
// 	}
// 	return _generateFile(&generateOptions{
// 		path:          path + "/protos/services/",
// 		serviceName:   serviceName,
// 		fileExtension: ".proto",
// 		fileTemplate:  "protoService.txt",
// 		data: struct {
// 			ServiceName      string
// 			UpperServiceName string
// 		}{
// 			ServiceName:      serviceName,
// 			UpperServiceName: inflect.Titleize(serviceName),
// 		},
// 	})
// }

// func _generateFile(options *generateOptions) error {
// 	file, err := _createFile(options)
// 	if err != nil {
// 		return err
// 	}
// 	return _writeTemplateContent(file, options)
// }

// func _createFile(options *generateOptions) (*os.File, error) {
// 	err := os.MkdirAll(options.path, os.ModePerm)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return os.Create(options.getFilePath())
// }

// func _writeTemplateContent(file *os.File, options *generateOptions) error {
// 	defer file.Close()
// 	tmpl := template.Must(template.ParseFiles(
// 		GoConfig.GetConfigStringValue("templatesPath") + options.fileTemplate),
// 	)
// 	return tmpl.Execute(file, options.data)
// }
