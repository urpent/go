package logx

import (
	"context"
	"strings"
	"testing"

	"github.com/urpent/go/ut"
)

func TestLoggerLocal(t *testing.T) {
	t.Run("Test Info", func(t *testing.T) {
		buf := &strings.Builder{}
		localStdout = buf
		logger := NewLoggerLocal()
		logger.Info(context.TODO(), "HelloWorld")
		ut.AssertEqual(t, "HelloWorld\n", strings.Split(buf.String(), " ")[3])
	})

	t.Run("Test Warn", func(t *testing.T) {
		buf := &strings.Builder{}
		localStdout = buf
		logger := NewLoggerLocal()
		logger.Warn(context.TODO(), "HelloWorld")
		ut.AssertEqual(t, "HelloWorld\n", strings.Split(buf.String(), " ")[3])
	})

	t.Run("Test Error", func(t *testing.T) {
		buf := &strings.Builder{}
		localStdout = buf
		logger := NewLoggerLocal()
		logger.Error(context.TODO(), "HelloWorld")
		ut.AssertEqual(t, "HelloWorld\n", strings.Split(buf.String(), " ")[3])
	})
}
