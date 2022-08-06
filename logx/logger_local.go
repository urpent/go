package logx

import (
	"context"
	"io"
	"log"
	"os"
)

const logFlags = log.LstdFlags

var localStdout io.Writer = os.Stdout

type LoggerLocal struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

func (l LoggerLocal) Info(_ context.Context, v ...interface{}) {
	l.infoLogger.Println(v...)
}

func (l LoggerLocal) Warn(_ context.Context, v ...interface{}) {
	l.warnLogger.Println(v...)
}

func (l LoggerLocal) Error(_ context.Context, v ...interface{}) {
	l.errorLogger.Println(v...)
}

func NewLoggerLocal() ExternalLogger {
	return LoggerLocal{
		infoLogger:  log.New(localStdout, "INFO: ", logFlags),
		warnLogger:  log.New(localStdout, "WARN: ", logFlags),
		errorLogger: log.New(localStdout, "ERROR: ", logFlags),
	}
}
