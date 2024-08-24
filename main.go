package main

import (
	log "github.com/sirupsen/logrus"
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
