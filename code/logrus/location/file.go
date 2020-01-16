package main

import (
	"os"
	"log"
	"github.com/sirupsen/logrus"
)

func main() {
	// You can have any number of instances.
	var consoleLog = logrus.New()
	consoleLog.Out = os.Stdout

	var fileLog = logrus.New()
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to log to file, using default stderr")
	}
	fileLog.Out = file
	defer file.Close()

	consoleLog.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	fileLog.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}
