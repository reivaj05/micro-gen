package generator

import (
	"fmt"
	"os"

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

func Generate(flags map[string]string, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("You didn't pass a name for the microservice")
	}
	// TODO: Check language is allowed
	language := flags["lang"]
	serviceName := args[0]
	if err := generators[language](serviceName); err != nil {
		rollback(serviceName)
		return err
	}
	return nil
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

func rollback(serviceName string) {
	os.RemoveAll(fmt.Sprintf("./%s", serviceName))
}
