package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/utrack/go-simple-chat/logger"
)

func loggerFunc(level logger.Level, format string, opts ...interface{}) {
	switch level {
	case logger.LevelDebug:
		logrus.Debugf(format, opts...)
	case logger.LevelWarn:
		logrus.Warnf(format, opts...)
	case logger.LevelError:
		logrus.Errorf(format, opts...)
	case logger.LevelFatal:
		logrus.Fatalf(format, opts...)
	default:
		logrus.Infof(format, opts...)
	}
}
