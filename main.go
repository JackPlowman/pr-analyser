package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	initLogging()
	RunHello()
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

func GitHubActionSummary(){
  action := os.Getenv("GITHUB_ACTION")
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
