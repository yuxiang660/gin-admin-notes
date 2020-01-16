package main

import (
	mongohook "github.com/LyricTian/logrus-mongo-hook"
	"github.com/sirupsen/logrus"
)

func main() {
	// mongodb url, database name, collection name
	hook := mongohook.DefaultWithURL("127.0.0.1:27017", "test", "t_log")
	defer hook.Flush()

	log := logrus.New()
	log.AddHook(hook)

	log.SetLevel(logrus.WarnLevel)
	// won't log to console and database because of info level
	log.WithField("foo-info", "bar-info").Infof("test info")
	// log on both console and database
	log.WithField("foo-warn", "bar-wan").Warn("test warn")
}