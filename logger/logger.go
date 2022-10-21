package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLog() {
	Logger = logrus.New()

}
func getLogger() *logrus.Entry {
	if Logger == nil {
		panic("logger is not init")
	}
	return Logger.WithContext(context.WithValue(context.Background(), "sendToFile", true))
}

func LogErrf(format string, v ...interface{}) {
	getLogger().Errorf(format, v)
}

func LogWarnf(format string, v ...interface{}) {
	getLogger().Warnf(format, v)
}

func LogInfof(format string, v ...interface{}) {
	getLogger().Infof(format, v)
}

func LogInfo(format string) {
	getLogger().Info(format)
}

func LogTraceJsonf(format string, v interface{}) {
	d, _ := json.Marshal(v)
	s := fmt.Sprintf(format, string(d))
	getLogger().Tracef(format, s)
}

func LogTracef(format string, v ...interface{}) {
	getLogger().Tracef(format, v)
}

func LogFatalf(format string, v ...interface{}) {
	getLogger().Fatalf(format, v)
}
