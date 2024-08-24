package main

import (
	hello "github.com/JackPlowman/github-pr-analyser/internal/hello"
	log "github.com/sirupsen/logrus"
)

func main() {
	initLogging()
	hello.RunHello()
}

// Init logging configuration
func initLogging() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}
