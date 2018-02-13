package operationsBuilder

import (
	"fmt"
	"strings"

	"github.com/reivaj05/GoJSON"

	"github.com/reivaj05/micro-gen/generator/utils"
)

func Build(opName, services string) error {
	fmt.Println("TODO: Implement operations")
	// return nil
	if err := createDirectories(opName); err != nil {
		return err
	}
	return createService(opName, strings.Split(services, ","))
}

func createDirectories(opName string) error {
	paths := []string{fmt.Sprintf("./%s", opName)}
	for _, path := range paths {
		if err := utils.CreateDir(path); err != nil {
			return err
		}
	}
	return nil
}

func createService(opName string, services []string) error {
	services = filterServices(services)
	return generateAllFiles(opName, services)
}

func filterServices(services []string) []string {
	// if docker, err := dockerWrapper.NewDockerRegistryManager(); err == nil {
	// 	reposResponse, err := docker.SearchRepos()
	// 	if err == nil {
	// 		services = filterAgainstDockerRegistryRepos(services, reposResponse)
	// 	}
	// }
	return services
}

func filterAgainstDockerRegistryRepos(
	services []string, reposResponse *GoJSON.JSONWrapper) (filteredServices []string) {

	repos := reposResponse.GetArrayFromPath("results")
	for _, service := range services {
		if serviceIsInDockerRegistry(repos, service) {
			filteredServices = append(filteredServices, service)
		}
	}
	return filteredServices
}

func serviceIsInDockerRegistry(repos []*GoJSON.JSONWrapper, service string) bool {
	for _, repo := range repos {
		if repo.HasPath("name") {
			if name, ok := repo.GetStringFromPath("name"); ok && name == service {
				return true
			}
		}
	}
	return false
}

func generateAllFiles(opName string, services []string) error {
	for _, service := range services {
		if err := generateFilesForService(opName, service); err != nil {
			return err
		}
	}
	return nil
}

func generateFilesForService(opName, service string) error {
	if err := utils.CreateDir(fmt.Sprintf("./%s/%s", opName, service)); err != nil {
		return err
	}
	if err := generateDeploymentFile(opName, service); err != nil {
		return err
	}
	return generateServiceFile(opName, service)
}

func generateDeploymentFile(opName, service string) error {
	options := utils.CreateFileOptions("deployment.yml", fmt.Sprintf("%s/", service),
		"deployment.gen", "", "operations", true)
	return utils.GenerateFile(opName, options)
}

func generateServiceFile(opName, service string) error {
	options := utils.CreateFileOptions("service.yml", fmt.Sprintf("%s/", service),
		"service.gen", "", "operations", true)
	return utils.GenerateFile(opName, options)
}
