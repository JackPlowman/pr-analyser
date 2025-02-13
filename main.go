package main

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func main() {
	initLogging()
	RunHello()
	CreateTempFolders()
	GitHubActionSummary()
}

// Init logging configuration
func initLogging() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func RunHello() {
	log.Info("Hello World")
}

// Generate GitHub Action Summary
func GitHubActionSummary() {
	action := os.Getenv("RUNNING_IN_GITHUB_ACTION")
	if action == "true" {
		log.Info("Running in GitHub Action, Generating Summary")
		gitHubActionSummaryFile := os.Getenv("GITHUB_STEP_SUMMARY")
		content := []byte("# Hello World")
		err := os.WriteFile(gitHubActionSummaryFile, content, 0600)
		if err != nil {
			panic(err)
		}
		log.Info("Summary Generated")
	}
}

// Create a temp folder with two subfolders: one for the repo default branch and one for the commit being reviewed
func CreateTempFolders() {
	tempDir, err := os.MkdirTemp("", "github-pr-analyser")
	if err != nil {
		log.Fatal("Failed to create temp directory: ", err)
	}

	defaultBranchDir := filepath.Join(tempDir, "default-branch")
	commitDir := filepath.Join(tempDir, "commit")

	err = os.Mkdir(defaultBranchDir, 0755)
	if err != nil {
		log.Fatal("Failed to create default branch directory: ", err)
	}

	err = os.Mkdir(commitDir, 0755)
	if err != nil {
		log.Fatal("Failed to create commit directory: ", err)
	}

	log.Infof("Created temp directories: %s, %s", defaultBranchDir, commitDir)
}
