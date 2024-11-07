package logger

import (
	"context"
	legolog "github.com/go-acme/lego/v4/log"
	fiberlog "github.com/gofiber/fiber/v2/log"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var _ fiberlog.AllLogger = (*logrusLogger)(nil)
var logger = logrus.New()

type logrusLogger struct {
	logger *logrus.Logger
}

func (c logrusLogger) Trace(v ...interface{}) {
	c.logger.Trace(v...)
}

func Trace(v ...interface{}) {
	logger.Trace(v...)
}

func (c logrusLogger) Debug(v ...interface{}) {
	c.logger.Debug(v...)
}

func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func (c logrusLogger) Info(v ...interface{}) {
	c.logger.Info(v...)
}

func Info(v ...interface{}) {
	logger.Info(v...)
}

func (c logrusLogger) Warn(v ...interface{}) {
	c.logger.Warn(v...)
}

func Warn(v ...interface{}) {
	logger.Warn(v...)
}

func (c logrusLogger) Error(v ...interface{}) {
	c.logger.Error(v...)
}

func Error(v ...interface{}) {
	logger.Error(v...)
}

func (c logrusLogger) Fatal(v ...interface{}) {
	c.logger.Fatal(v...)
}

func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

func (c logrusLogger) Panic(v ...interface{}) {
	c.logger.Panic(v...)
}

func Panic(v ...interface{}) {
	logger.Panic(v...)
}

func (c logrusLogger) Tracef(format string, v ...interface{}) {
	c.logger.Tracef(format, v...)
}

func Tracef(format string, v ...interface{}) {
	logger.Tracef(format, v...)
}

func (c logrusLogger) Debugf(format string, v ...interface{}) {
	c.logger.Debugf(format, v...)
}

func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

func (c logrusLogger) Infof(format string, v ...interface{}) {
	c.logger.Infof(format, v...)
}

func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

func (c logrusLogger) Warnf(format string, v ...interface{}) {
	c.logger.Warnf(format, v...)
}

func Warnf(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}

func (c logrusLogger) Errorf(format string, v ...interface{}) {
	c.logger.Errorf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

func (c logrusLogger) Fatalf(format string, v ...interface{}) {
	c.logger.Fatalf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}

func (c logrusLogger) Panicf(format string, v ...interface{}) {
	c.logger.Panicf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	logger.Panicf(format, v...)
}

func (c logrusLogger) logw(level logrus.Level, msg string, keysAndValues ...interface{}) {
	fields := logrus.Fields{}
	for i := 0; i < len(keysAndValues); i += 2 {
		fields[keysAndValues[i].(string)] = keysAndValues[i+1]
	}
	c.logger.WithFields(fields).Log(level, msg)
}

func (c logrusLogger) Tracew(msg string, keysAndValues ...interface{}) {
	c.logw(logrus.TraceLevel, msg, keysAndValues...)
}

func (c logrusLogger) Debugw(msg string, keysAndValues ...interface{}) {
	c.logw(logrus.DebugLevel, msg, keysAndValues...)
}

func (c logrusLogger) Infow(msg string, keysAndValues ...interface{}) {
	c.logw(logrus.InfoLevel, msg, keysAndValues...)
}

func (c logrusLogger) Warnw(msg string, keysAndValues ...interface{}) {
	c.logw(logrus.WarnLevel, msg, keysAndValues...)
}

func (c logrusLogger) Errorw(msg string, keysAndValues ...interface{}) {
	c.logw(logrus.ErrorLevel, msg, keysAndValues...)
}

func (c logrusLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	c.logw(logrus.FatalLevel, msg, keysAndValues...)
}

func (c logrusLogger) Panicw(msg string, keysAndValues ...interface{}) {
	c.logw(logrus.PanicLevel, msg, keysAndValues...)
}

func (c logrusLogger) SetLevel(level fiberlog.Level) {
	logger.SetLevel(logrus.Level(level))
}

func (c logrusLogger) SetOutput(writer io.Writer) {
	logger.SetOutput(writer)
}

func (c logrusLogger) WithContext(ctx context.Context) fiberlog.CommonLogger {
	entry := c.logger.WithContext(ctx)
	return &logrusLogger{
		logger: entry.Logger,
	}
}

func init() {
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	logger.SetOutput(os.Stdout)
	fiberlog.SetLogger(&logrusLogger{
		logger: logger,
	})
	legolog.Logger = logger
}
