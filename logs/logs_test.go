package logs

import (
	"context"
	"strings"
	"testing"

	"github.com/urpent/go/ut"
)

func TestInfo(t *testing.T) {
	t.Run("Test Info", func(t *testing.T) {
		buf := &strings.Builder{}
		localStdout = buf
		logger := NewLoggerLocal()
		SetLogger(logger)
		Info("HelloWorld")
		s := strings.Split(buf.String(), "\n")[0]
		ut.AssertEqual(t, "HelloWorld", strings.Split(s, " ")[3])
	})

	t.Run("Test Warn", func(t *testing.T) {
		buf := &strings.Builder{}
		localStdout = buf
		logger := NewLoggerLocal()
		SetLogger(logger)
		Warn("HelloWorld")
		s := strings.Split(buf.String(), "\n")[0]
		ut.AssertEqual(t, "HelloWorld", strings.Split(s, " ")[3])
	})

	t.Run("Test Error", func(t *testing.T) {
		buf := &strings.Builder{}
		localStdout = buf
		logger := NewLoggerLocal()
		SetLogger(logger)
		Error("HelloWorld")
		s := strings.Split(buf.String(), "\n")[0]
		ut.AssertEqual(t, "HelloWorld", strings.Split(s, " ")[3])
	})

	t.Run("Test Error with info", func(t *testing.T) {
		buf := &strings.Builder{}
		localStdout = buf
		logger := NewLoggerLocal()
		SetLogger(logger)
		Error("HelloWorld", Field{
			Name: "request",
			Data: struct {
				ID int
			}{1},
		})
		s := strings.Split(buf.String(), "\n")
		ut.AssertEqual(t, "HelloWorld", strings.Split(s[0], " ")[3])
		ut.AssertEqual(t, ` "Name": "request",`, s[3])
		ut.AssertEqual(t, `  "ID": 1`, s[5])
	})

	t.Run("Test Error with context", func(t *testing.T) {
		buf := &strings.Builder{}
		localStdout = buf
		logger := NewLoggerLocal()
		SetLogger(logger)
		Error("HelloWorld", context.TODO(), Field{
			Name: "request",
			Data: struct {
				ID int
			}{1},
		})
		s := strings.Split(buf.String(), "\n")
		ut.AssertEqual(t, "HelloWorld", strings.Split(s[0], " ")[3])
		ut.AssertEqual(t, ` "Name": "request",`, s[3])
		ut.AssertEqual(t, `  "ID": 1`, s[5])
	})
}

func TestString(t *testing.T) {
	t.Run("Test Simple String", func(t *testing.T) {
		result := jsonString(123)
		ut.AssertEqual(t, `123`, result)
	})

	t.Run("Test Simple String", func(t *testing.T) {
		obj := struct {
			ID string
		}{
			ID: "abc",
		}
		result := jsonString(obj)
		ut.AssertEqual(t, "{\n \"ID\": \"abc\"\n}", result)
	})
}
