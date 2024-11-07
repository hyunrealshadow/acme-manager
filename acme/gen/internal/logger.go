package internal

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger = &logrus.Logger{
	Out:       os.Stdout,
	Formatter: &logrus.TextFormatter{},
	Level:     logrus.InfoLevel,
	ExitFunc:  os.Exit,
}
