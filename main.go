package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	initLogging()
	RunHello()
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
func GitHubActionSummary(){
  action := os.Getenv("RUNNING_IN_GITHUB_ACTION")
	if action == "true" {
		log.Info("Running in GitHub Action, Generating Summary")
		gitHubActionSummaryFile := os.Getenv("GITHUB_STEP_SUMMARY")
		content := []byte("# Hello World")
		err := os.WriteFile(gitHubActionSummaryFile, content, 0644)
		if err != nil {
			panic(err)
	}
		log.Info("Summary Generated")
	}
}
