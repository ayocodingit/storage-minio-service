package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func New() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{"msg": "message", "time": "timestamp"},
	})
	log.SetOutput(os.Stdout)
	return log
}
