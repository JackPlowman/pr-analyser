package main

import (
	log "github.com/sirupsen/logrus"
)
import (
	hello "github.com/JackPlowman/github-pr-analyser/src/hello"
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
