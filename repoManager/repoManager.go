package repoManager

import (
	"fmt"
)

type repoCreator func() error

var repoProviders = map[string]repoCreator{
	"github":		createGitHubRepo,
	"bitbucket":    createBitBucketRepo,
	"gitlab":       createGitLabRepo,
}


func CreateRepo(provider string) error {
	return repoProviders[provider]()
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