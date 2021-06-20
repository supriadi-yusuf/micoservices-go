package option_a

import (
	"microservices-go/config"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	Log *logrus.Logger
)

func init() {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		level = logrus.DebugLevel
	}

	Log = &logrus.Logger{
		Level:     level,
		Out:       os.Stdout,
		Formatter: &logrus.JSONFormatter{},
	}

}

func Info(msg string, tags ...string) {
	Log.WithFields(ParseFields(tags...)).Info(msg)
}

func ParseFields(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))

	for _, tag := range tags {
		elms := strings.Split(tag, ":")
		result[strings.TrimSpace(elms[0])] = strings.TrimSpace(elms[1])
	}

	return result
}
