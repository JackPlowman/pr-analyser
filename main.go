package main

import (
	"os"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

func main() {
	initLogging()
	RunHello()
	CheckGitHubWorkspace()
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

// Check if github/workspace folder has files
func CheckGitHubWorkspace() {
	files, err := ioutil.ReadDir("/github/workspace")
	if err != nil {
		log.Error("Error reading github/workspace folder: ", err)
		return
	}
	if len(files) == 0 {
		log.Error("Repository not checked out")
	} else {
		log.Debug("Files found in github/workspace folder")
	}
}
