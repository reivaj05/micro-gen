package generator

import (
	"fmt"
	"os"

	goBuilder "github.com/reivaj05/micro-gen/generator/builders/go"
	jsBuilder "github.com/reivaj05/micro-gen/generator/builders/javascript"
	managerBuilder "github.com/reivaj05/micro-gen/generator/builders/manager"
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

func GenerateService(flags map[string]string, args ...string) error {
	language := flags["lang"]
	if err := validateGenerateServiceParameters(language, args...); err != nil {
		return err
	}
	serviceName := args[0]
	if err := generators[language](serviceName); err != nil {
		rollback(serviceName)
		return err
	}
	return nil
}

func validateGenerateServiceParameters(language string, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("You didn't pass a name for the microservice")
	}
	if _, ok := generators[language]; !ok {
		return fmt.Errorf("Programming language not allowed")
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

func GenerateManager(flags map[string]string, args ...string) error {
	fmt.Println("TODO: Generate manager")
	if err := managerBuilder.Build("manager"); err != nil {
		rollback("manager")
		return err
	}
	return nil
}

func rollback(serviceName string) {
	os.RemoveAll(fmt.Sprintf("./%s", serviceName))
}
