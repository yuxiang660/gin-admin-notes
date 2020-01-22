package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"runtime"
	"strings"
)

var logger = logrus.New()

// Fields wraps logrus.Fields, which is a map[string]interface{}
type Fields logrus.Fields

func SetLogLevel(level logrus.Level) {
	logger.Level = level
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Error(args...)
	}
}

// Debug logs a message with fields at level Debug on the standard logger.
func ErrorWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Error(l)
	}
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func main() {
	fmt.Println("logrus to log line number of the file")

	SetLogLevel(logrus.ErrorLevel)

	Error(1, "Message, to")
	ErrorWithFields(Fields{"animal": "walrus"}, Fields{"animal": "walrus"})
}
