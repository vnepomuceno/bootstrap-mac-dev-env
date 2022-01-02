package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"gopkg.in/src-d/go-git.v4"
	"log"
	"os"
	"strings"
)

func makeDir(directory string) {
	log.Println(fmt.Sprintf("Creating directory '%s'", directory))
	_, err := os.Stat(directory)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(directory, 0755)
		if errDir != nil {
			log.Fatal(err)
		} else {
			log.Println(fmt.Sprintf("Successfully created directory '%s'", directory))
		}
	} else {
		log.Println("Directory already exists and will not be recreated")
	}
}

func listGithubRepos() {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "... Your Token ..."},
	)
	client := github.NewClient(oauth2.NewClient(context.Background(), tokenSource))

	repos, _, _ := client.Repositories.List(context.Background(), "", nil)
	log.Println(fmt.Sprintf("Listed repos %s", repos))
}

func cloneGitRepository(githubUrl string, workspaceDir string) {
	repoName := strings.Split(strings.Split(githubUrl, "/")[4], ".")[0]
	repoDir := workspaceDir + "/" + repoName
	makeDir(repoDir)

	repo, err := git.PlainClone(repoDir, false, &git.CloneOptions{
		URL:               githubUrl,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Sprintf("Successfully cloned repository from %s", repo))
}

func main() {
	// Sets arguments to be received by shell command
	workspaceDir := "/Users/valternepomuceno/go-workspace"
	githubUrl := "https://github.com/vnepomuceno/gmail-fisher.git"

	// Clones all repos into workspace directory
	listGithubRepos()
	makeDir(workspaceDir)
	cloneGitRepository(githubUrl, workspaceDir)
}
