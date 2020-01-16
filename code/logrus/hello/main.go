package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	// output: time="2020-01-16T19:16:03+08:00" level=info msg="A walrus appears" animal=walrus
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
