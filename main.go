package main

import (
	"fmt"
	"os"

	"github.com/reivaj05/micro-gen/generator"

	"github.com/reivaj05/GoCLI"
	"github.com/reivaj05/GoConfig"
	"github.com/reivaj05/GoLogger"
)

const appName = "micro-gen"

func main() {
	setup()
	startApp()
}

func setup() {
	startConfig()
	startLogger()
}

func startConfig() {
	if err := GoConfig.Init(createConfigOptions()); err != nil {
		finishExecution("Error while starting config", map[string]interface{}{
			"error": err.Error(),
		})
	}
}

func createConfigOptions() *GoConfig.ConfigOptions {
	return &GoConfig.ConfigOptions{
		ConfigType: "json",
		ConfigFile: "config",
		ConfigPath: fmt.Sprintf("%s/src/github.com/reivaj05/micro-gen/", os.Getenv("GOPATH")),
	}
}

func startLogger() {
	if err := GoLogger.Init(createLoggerOptions()); err != nil {
		finishExecution("Error while loading logger", map[string]interface{}{
			"error": err.Error(),
		})
	}
}

func createLoggerOptions() *GoLogger.LoggerOptions {
	return &GoLogger.LoggerOptions{
		OutputFile: fmt.Sprintf("%s-log.json", appName),
		Path:       "log/",
		LogLevel:   getLogLevel(),
	}
}

func getLogLevel() int {
	levels := map[string]int{"DEBUG": GoLogger.DEBUG, "INFO": GoLogger.INFO,
		"WARNING": GoLogger.WARNING, "ERROR": GoLogger.ERROR,
		"PANIC": GoLogger.PANIC, "FATAL": GoLogger.FATAL,
	}
	if level, ok := levels[GoConfig.GetConfigStringValue("logLevel")]; ok {
		return level
	}
	return GoLogger.INFO
}

func startApp() {
	if err := GoCLI.StartCLI(createCLIOptions()); err != nil {
		finishExecution("Error while starting application", map[string]interface{}{
			"error": err.Error(),
		})
	}
}

func createCLIOptions() *GoCLI.Options {
	return &GoCLI.Options{
		AppName:  appName,
		AppUsage: "Tool to create different services in different languages and tooling to handle those services",
		Commands: createCommands(),
	}
}

func createCommands() []*GoCLI.Command {
	return []*GoCLI.Command{
		createServiceCommand(),
		createToolingCommand(),
		createOperationsCommand(),
	}
}

func createServiceCommand() *GoCLI.Command {
	return &GoCLI.Command{
		Name:        "create-service",
		Usage:       "Create a new service project in the language of your preference",
		Action:      generator.GenerateService,
		StringFlags: getCreateServiceStringFlags(),
	}
}

func getCreateServiceStringFlags() []*GoCLI.StringFlag {
	return []*GoCLI.StringFlag{
		createStringFlag("lang", "Language of the microservice to be created", "go"),
		createStringFlag("repo-provider", "Service to handle repos(github, gitlab)", "github"),
		createStringFlag("ci-provider", "Service to handle CI integration(travis)", "travis"),
	}
}

func createToolingCommand() *GoCLI.Command {
	return &GoCLI.Command{
		Name:        "create-tooling",
		Usage:       "Create a new tooling project to handle the services you previously created",
		Action:      generator.GenerateTooling,
		StringFlags: getCreateToolingStringFlags(),
	}
}

func createOperationsCommand() *GoCLI.Command {
	return &GoCLI.Command{
		Name:        "create-operations",
		Usage:       "Create a new kubernetes project to handle the services you previously created",
		Action:      generator.GenerateOperations,
		StringFlags: getCreateToolingStringFlags(),
	}
}

func getCreateToolingStringFlags() []*GoCLI.StringFlag {
	return []*GoCLI.StringFlag{
		createStringFlag("services", "Space separated list of the services you want to manage", ""),
		createStringFlag("repo-provider", "Service to handle repos(github, gitlab)", "github"),
	}
}

func createStringFlag(name, usage, _default string) *GoCLI.StringFlag {
	return &GoCLI.StringFlag{
		Name:    name,
		Usage:   usage,
		Default: _default,
	}
}

func finishExecution(msg string, fields map[string]interface{}) {
	GoLogger.LogFatal(msg, fields)
}
