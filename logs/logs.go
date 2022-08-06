// Package logs is a simple logging lib that shows calling func line, info and provide context for tracing.
package logs

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
)

// Logger not using struct because Go does not support generic methods.
var externalLoggers = []ExternalLogger{NewLoggerLocal()}

type ExternalLogger interface {
	Info(ctx context.Context, v ...any)
	Warn(ctx context.Context, v ...any)
	Error(ctx context.Context, v ...any)
}

// SetLogger will set external loggers. This should be set ONCE only.
func SetLogger(ls ...ExternalLogger) {
	if ls != nil {
		externalLoggers = ls
	}
}

type Field struct {
	Name string
	Data any
}

// Info will log and return first parameter p
// Example 1: logs.Info(err)
// Example 2 with extra request info: logs.Error(err, logs.Field{Name: "myRequest", Data: req})
// Example 3: logs.Info(errx.Wrap(errs.ErrUnauthorized, errs.ErrUserNotFound, err), logs.Field{"myRequest", req})
func Info[T any](p T, extraInfos ...any) T {
	for _, lg := range externalLoggers {
		logF(lg.Info, p, extraInfos...)
	}

	return p
}

// Warn will log and return first parameter p
// Example: logs.Warn(errx.Wrap(errs.ErrUnauthorized, errs.ErrUserNotFound, err), logs.Field{"myRequest", req})
func Warn[T any](p T, extraInfos ...any) T {
	for _, lg := range externalLoggers {
		logF(lg.Warn, p, extraInfos...)
	}

	return p
}

// Error will log and return first parameter p1
// Example 1: logs.Error(err)
// Example 2 with extra request info: logs.Error(err, logs.Field{"myRequest", Data: req})
// Example 3: logs.Error(errx.Wrap(errs.ErrUnauthorized, errs.ErrUserNotFound, err))
func Error[T any](p T, extraInfos ...any) T {
	for _, lg := range externalLoggers {
		logF(lg.Error, p, extraInfos...)
	}

	return p
}

func jsonString[T any](obj T) string {
	a, _ := json.MarshalIndent(obj, "", " ")
	return string(a)
}

// logF will format the info and execute log functionality
func logF(logFunc func(ctx context.Context, v ...interface{}), v any, infos ...any) {
	all := []any{v, "\n", trace4()}

	var ctx context.Context // Optional

	for _, info := range infos {
		switch p := info.(type) {
		case context.Context:
			ctx = p
		default:
			all = append(all, "\n"+jsonString(info))
		}
	}

	logFunc(ctx, all...)
}

func trace4() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(4, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function)
}
