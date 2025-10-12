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
	if payload != nil {
		Log.WithFields(logrus.Fields{
			"event":   event,
			"payload": payload,
		}).Info("Event Info")
	} else {
		Log.WithFields(logrus.Fields{
			"event": event,
		}).Info("Event Info")
	}
}

func Error(event string, err error, payload any) {
	fields := logrus.Fields{
		"event": event,
		"error": err.Error(),
	}
	if payload != nil {
		fields["payload"] = payload
	}
	Log.WithFields(fields).Error("Event Error")
}
