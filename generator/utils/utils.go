package utils

import "os"

type GenerateOptions struct {
	ServiceName   string
	FileExtension string
	FileTemplate  string
	FileName      string
	Data          interface{}
}

func GetMicroGenPath() string {
	const relativePath = "/src/github.com/reivaj05/micro-gen"
	goPath := os.Getenv("GOPATH")
	return goPath + relativePath
}
