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
		AppUsage: "TODO: Set app usage",
		Commands: createCommands(),
		// StringFlags: createStringFlags(),
		// DefaultAction: server.Serve,
	}
}

func createCommands() []*GoCLI.Command {
	return []*GoCLI.Command{
		&GoCLI.Command{
			Name:   "create-service",
			Usage:  "TODO: Set create-service usage",
			Action: generator.GenerateService,
			StringFlags: getCreateServiceStringFlags(),
		},
		&GoCLI.Command{
			Name:   "create-manager",
			Usage:  "TODO: Set create-manager usage",
			Action: generator.GenerateManager,
			StringFlags: getCreateManagerStringFlags(),
		},
	}
}

func getCreateServiceStringFlags() []*GoCLI.StringFlag {
	return []*GoCLI.StringFlag{
		&GoCLI.StringFlag{
			Name:   "lang",
			Usage:  "Language of the microservice to be created",
			Default: "go",
		},
	}
}

func getCreateManagerStringFlags() []*GoCLI.StringFlag {
	return []*GoCLI.StringFlag{
		&GoCLI.StringFlag{
			Name:   "services",
			Usage:  "Space separated list of the services you want to manage. TODO",
		},
	}
}

func finishExecution(msg string, fields map[string]interface{}) {
	GoLogger.LogFatal(msg, fields)
}
