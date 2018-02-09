package generator

import (
	"fmt"
	"os"

	"github.com/reivaj05/micro-gen/ci-manager"
	goBuilder "github.com/reivaj05/micro-gen/generator/builders/go"
	jsBuilder "github.com/reivaj05/micro-gen/generator/builders/javascript"
	operationsBuilder "github.com/reivaj05/micro-gen/generator/builders/operations"
	pythonBuilder "github.com/reivaj05/micro-gen/generator/builders/python"
	rubyBuilder "github.com/reivaj05/micro-gen/generator/builders/ruby"
	rustBuilder "github.com/reivaj05/micro-gen/generator/builders/rust"
	toolingBuilder "github.com/reivaj05/micro-gen/generator/builders/tooling"
	"github.com/reivaj05/micro-gen/repo-manager"
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
	if err := repoManager.CreateRepo(serviceName, flags["repo-provider"]); err != nil {
		return err
	}
	return CIManager.ConnectWithCIProvider(serviceName, flags["ci-provider"])
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

func GenerateTooling(flags map[string]string, args ...string) error {
	services := flags["services"]
	if err := toolingBuilder.Build("tooling", services); err != nil {
		rollback("tooling")
		return err
	}
	return repoManager.CreateRepo("tooling", flags["repo-provider"])
}

func GenerateOperations(flags map[string]string, args ...string) error {
	services := flags["services"]
	if err := operationsBuilder.Build("operations", services); err != nil {
		rollback("operations")
		return err
	}
	return repoManager.CreateRepo("operations", flags["repo-provider"])
}

func rollback(serviceName string) {
	os.RemoveAll(fmt.Sprintf("./%s", serviceName))
}
