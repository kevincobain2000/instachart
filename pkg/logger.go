package pkg

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	once sync.Once
	log  *logrus.Logger
)

const (
	DEFAULT_LOG_FILE = "./logs/go-chart-to-picture.log"
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
		// set logs to file and log rotate
		// Set up lumberjack log file rotation
		log.SetOutput(&lumberjack.Logger{
			Filename:   DEFAULT_LOG_FILE, // specify your log file path
			MaxSize:    10,               // max size in megabytes
			MaxBackups: 5,                // max number of old log files to keep
			MaxAge:     28,               // max age in days to keep a log file
			Compress:   true,             // whether to compress log files
		})
	})
	return log
}
