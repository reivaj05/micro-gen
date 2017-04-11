package generator

import (
	goBuilder "github.com/reivaj05/micro-gen/generator/builders/go"
	jsBuilder "github.com/reivaj05/micro-gen/generator/builders/javascript"
	pythonBuilder "github.com/reivaj05/micro-gen/generator/builders/python"
	rubyBuilder "github.com/reivaj05/micro-gen/generator/builders/ruby"
	rustBuilder "github.com/reivaj05/micro-gen/generator/builders/rust"
)

type generator func() error

var generators = map[string]generator{
	"go":         generateGo,
	"python":     generatePython,
	"ruby":       generateRuby,
	"javascript": generateJS,
	"rust":       generateRust,
}

func Generate(args ...string) error {
	language := args[0]
	return generators[language]()
}

func generateGo() error {
	return goBuilder.Build()
}

func generateJS() error {
	return jsBuilder.Build()
}

func generatePython() error {
	return pythonBuilder.Build()
}

func generateRuby() error {
	return rubyBuilder.Build()
}

func generateRust() error {
	return rustBuilder.Build()
}

func rollback() error {
	return nil
}
