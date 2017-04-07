package generator

import (
	goGen "github.com/reivaj05/micro-gen/generator/generators/go"
	jsGen "github.com/reivaj05/micro-gen/generator/generators/javascript"
	pythonGen "github.com/reivaj05/micro-gen/generator/generators/python"
	rubyGen "github.com/reivaj05/micro-gen/generator/generators/ruby"
	rustGen "github.com/reivaj05/micro-gen/generator/generators/rust"
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
	return goGen.Build()
}

func generateJS() error {
	return jsGen.Build()
}

func generatePython() error {
	return pythonGen.Build()
}

func generateRuby() error {
	return rubyGen.Build()
}

func generateRust() error {
	return rustGen.Build()
}

func rollback() error {
	return nil
}
