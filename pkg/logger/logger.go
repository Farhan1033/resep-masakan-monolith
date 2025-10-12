package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func Init() {
	Log = logrus.New()
	Log.SetOutput(os.Stdout)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Log.SetLevel(logrus.InfoLevel)
}

func Info(event string, payload any) {
	Log.WithFields(logrus.Fields{"event": event}).Info("Event Info")
}

func Error(event string, err error, payload any) {
	Log.WithFields(logrus.Fields{"event": event}).Error(err.Error())
}
