// Package ut provides commonly used func for unit testing
package ut

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/qdm12/reprint"
)

// AssertEqual return result of comparison. AssetEqual is wrapping google compare which will try to use Equal func of the object first.
// Don't use reflect.DeepEqual directly because it will not compare time correctly. Recommended to use AssertEqual
//
// Note: If comparing struct that contain unexported field, must add cmp.AllowUnexported(expected, actual) as opts.
// However, this option only allow comparing of value struct. Therefore, we need to dereference pointer struct.
func AssertEqual(t *testing.T, expected, actual interface{}, opts ...cmp.Option) bool {
	t.Helper()
	if errA, ok := expected.(error); ok {
		if errB, ok := actual.(error); ok {
			return assertEqualError(t, errA, errB)
		}
		return false
	}

	if !cmp.Equal(expected, actual, opts...) {
		t.Errorf("Not equal: \n"+
			"\nwant: %#v;"+
			"\ngot : %#v; "+
			"\ndiff: ---expected   +++actual\n"+
			" %s\n", expected, actual, cmp.Diff(expected, actual))
		return false
	}
	return true
}

func AssertNotEqual(t *testing.T, expected, actual interface{}, opts ...cmp.Option) bool {
	if errA, ok := expected.(error); ok {
		if errB, ok := actual.(error); ok {
			return assertEqualError(t, errA, errB)
		}
		return false
	}

	if cmp.Equal(expected, actual, opts...) {
		t.Helper()
		t.Errorf("Equal: \n"+
			"\ndon;t want: %#v;"+
			"\nbut got : %#v; "+
			"\ndiff: ---expected   +++actual\n"+
			" %s\n", expected, actual, cmp.Diff(expected, actual))
		return false
	}
	return true
}

// assertEqualError will compare error
func assertEqualError(t *testing.T, expected, actual error) bool {
	var expectedErr, actualErr interface{} = expected, actual

	if expected != nil {
		expectedErr = expected.Error()
	}

	if actual != nil {
		actualErr = actual.Error()
	}

	if !cmp.Equal(expectedErr, actualErr, cmpopts.EquateErrors()) {
		t.Helper()
		t.Errorf("Not equal: \n"+
			"\nwant: %#v;"+
			"\ngot : %#v; ", expectedErr, actualErr)
		return false
	}
	return true
}

// Clone return deep copy of a
// reprint lib is better than other deep copy lib due to its ability to copy unexported field too.
func Clone(a interface{}) interface{} {
	return reprint.This(a)
}

// TimeParse will parse time from string to *time.Time. Panic if parse failed
func TimeParse(layout, timeStr string) *time.Time {
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		panic(fmt.Sprintf("test case time layout is incorrect. %s is not %s", timeStr, layout))
	}

	return &parsedTime
}

// Ptr return the pointer of the value
func Ptr[T any](value T) *T {
	return &value
}
