package pkg

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	once sync.Once
	log  *logrus.Logger
)

const (
	DEFAULT_LOG_FILE = "./logs/instachart.log"
)

func Logger() *logrus.Logger {
	once.Do(func() {
		fmt.Println("Logger initialized")
		fmt.Println("Logging to file: ", DEFAULT_LOG_FILE)
		log = logrus.New()
		log.SetFormatter(&logrus.TextFormatter{
			DisableColors:   false,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
		log.SetReportCaller(false)
		log.SetLevel(logrus.InfoLevel)
	})
	return log
}
