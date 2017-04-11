package generator

import (
	goBuilder "github.com/reivaj05/micro-gen/generator/builders/go"
	jsBuilder "github.com/reivaj05/micro-gen/generator/builders/javascript"
	pythonBuilder "github.com/reivaj05/micro-gen/generator/builders/python"
	rubyBuilder "github.com/reivaj05/micro-gen/generator/builders/ruby"
	rustBuilder "github.com/reivaj05/micro-gen/generator/builders/rust"
)

type generator func(serviceName string) error

var generators = map[string]generator{
	"go":         generateGo,
	"python":     generatePython,
	"ruby":       generateRuby,
	"javascript": generateJS,
	"rust":       generateRust,
}

func Generate(args ...string) error {
	language := args[0]
	// TODO: Update service name from a
	serviceName := "serviceMock"
	return generators[language](serviceName)
}

func generateGo(serviceName string) error {
	return goBuilder.Build(serviceName)
}

func generateJS(serviceName string) error {
	return jsBuilder.Build(serviceName)
}

func generatePython(serviceName string) error {
	return pythonBuilder.Build(serviceName)
}

func generateRuby(serviceName string) error {
	return rubyBuilder.Build(serviceName)
}

func generateRust(serviceName string) error {
	return rustBuilder.Build(serviceName)
}

func rollback() error {
	return nil
}
