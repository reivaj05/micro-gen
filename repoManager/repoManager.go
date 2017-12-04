package repoManager

import (
	"fmt"
	"os/exec"
)

type repoCreator func() error

var repoProviders = map[string]repoCreator{
	"github":		createGitHubRepo,
	"bitbucket":    createBitBucketRepo,
	"gitlab":       createGitLabRepo,
}


func CreateRepo(serviceName, provider string) error {
	if err := createLocalRepo(serviceName); err != nil {
		return err
	}
	if creator, ok := repoProviders[provider]; ok {
		return creator()
	}
	return fmt.Errorf("Repo provider '%s' not supported", provider)
}

func createLocalRepo(serviceName string) error {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("cd %s; git init", serviceName))
	fmt.Println("Creating local repo...")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func createGitHubRepo() error {
	fmt.Println("TODO: Create github repo")
	return nil
}

func createBitBucketRepo() error {
	fmt.Println("TODO: Create bitbucket repo")
	return nil
}

func createGitLabRepo() error {
	fmt.Println("TODO: Create gitlab repo")
	return nil	
}