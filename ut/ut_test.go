package ut

import (
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestAssertEqual(t *testing.T) {
	timeNow := time.Now()

	tests := []struct {
		name       string
		expected   interface{}
		actual     interface{}
		wantResult bool
	}{
		{"OK, compare string", "123", "123", true},
		{"Not OK, string value and string pointer must return false", "123", Ptr("123"), false},
		{"Not OK, non-matching string value must return false", "123", "1234", false},
		{"OK, compare struct int", struct{ Num int }{Num: 123}, struct{ Num int }{Num: 123}, true},
		{"Not OK, struct int value and struct int pointer must return false", struct{ Num int }{Num: 123}, &struct{ Num int }{Num: 123}, false},
		{"Not OK, non-matching int value must return false", struct{ Num int }{Num: 123}, struct{ Num int }{Num: 1234}, false},
		{"OK, compare struct time", struct{ Time *time.Time }{Time: &timeNow}, struct{ Time *time.Time }{Time: TimeParse(time.RFC3339Nano, timeNow.Format(time.RFC3339Nano))}, true},
		{"Not OK, struct time value and struct time pointer must return false", struct{ Time *time.Time }{Time: &timeNow}, struct{ Time time.Time }{Time: *TimeParse(time.RFC3339Nano, timeNow.Format(time.RFC3339Nano))}, false},
		{"Not OK, non-matching time value must return false", struct{ Time *time.Time }{Time: &timeNow}, struct{ Time *time.Time }{Time: TimeParse(time.RFC3339Nano, timeNow.Add(time.Nanosecond).Format(time.RFC3339Nano))}, false},
	}

	mockT := new(testing.T)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := AssertEqual(mockT, tc.expected, tc.actual)
			if result != tc.wantResult {
				t.Errorf("AssertEqual(%#v, %#v) should return %#v, \ndiff: %s", tc.expected, tc.actual, tc.wantResult, cmp.Diff(tc.expected, tc.actual))
			}
		})
	}
}

func TestAssertEqual_error(t *testing.T) {
	err1 := errors.New("err1")

	tests := []struct {
		name       string
		expected   error
		actual     error
		wantResult bool
	}{
		{name: "OK, compare nil and nil", wantResult: true},
		{name: "Not OK, compare nil and error", expected: nil, actual: errors.New("actual error"), wantResult: false},
		{name: "Not OK, compare error and nil", expected: errors.New("expect error"), actual: nil, wantResult: false},
		{name: "OK, compare same error", expected: err1, actual: err1, wantResult: true},
		{name: "OK, compare same error message", expected: errors.New("error1"), actual: errors.New("error1"), wantResult: true},
		{name: "Not OK, compare different error", expected: errors.New("error1"), actual: errors.New("error2"), wantResult: false},
	}

	mockT := new(testing.T)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := AssertEqual(mockT, tc.expected, tc.actual)
			if tc.wantResult != result {
				t.Errorf("AssertEqual(%#v, %#v) should return %t, but got %t. \ndiff: %s", tc.expected, tc.actual, tc.wantResult, result, cmp.Diff(tc.expected, tc.actual, cmpopts.EquateErrors()))
			}
		})
	}
}

type testAssestEqualObj struct {
	Text        string
	IgnoredText string
}

// Equal is standard way to write comparing. It will be used by google cmp if it exist.
func (t testAssestEqualObj) Equal(o testAssestEqualObj) bool {
	return t.Text == o.Text
}

func TestAssertEqual_UsingEqualFunc(t *testing.T) {
	tests := []struct {
		name       string
		expected   *testAssestEqualObj
		actual     *testAssestEqualObj
		wantResult bool
	}{
		{"OK", &testAssestEqualObj{Text: "1", IgnoredText: "Xx"}, &testAssestEqualObj{Text: "1", IgnoredText: "12"}, true},
		{"Not OK, text is different", &testAssestEqualObj{Text: "2", IgnoredText: "abc"}, &testAssestEqualObj{Text: "1", IgnoredText: "efg"}, false},
	}

	mockT := new(testing.T)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := AssertEqual(mockT, tc.expected, tc.actual)
			if result != tc.wantResult {
				t.Errorf("AssertEqual(%#v, %#v) should return %#v, \ndiff: %s", tc.expected, tc.actual, tc.wantResult, cmp.Diff(tc.expected, tc.actual))
			}
		})
	}
}

func TestTimeParse(t *testing.T) {
	validTime, _ := time.Parse(time.RFC3339, "2020-01-04T14:30:28.152Z")

	tests := []struct {
		name       string
		timeLayout string
		timeString string
		wantResult *time.Time
	}{
		{name: "OK", timeLayout: time.RFC3339, timeString: "2020-01-04T14:30:28.152Z", wantResult: &validTime},
		{name: "Not Ok", timeLayout: time.RFC3339, timeString: "123456"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tc.wantResult != nil {
						t.Error("panic expected but did not")
					}
				}
			}()

			result := TimeParse(tc.timeLayout, tc.timeString)
			AssertEqual(t, tc.wantResult, result)
		})
	}
}

type testCloneObj struct {
	Text string
}

func TestClone(t *testing.T) {
	tests := []struct {
		name       string
		toClone    interface{}
		wantResult interface{}
	}{
		{"OK, clone string", "123", "123"},
		{"OK, clone struct", testCloneObj{Text: "567"}, testCloneObj{Text: "567"}},
		{"OK, clone pointer struct", &testCloneObj{Text: "567"}, &testCloneObj{Text: "567"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Clone(tc.toClone)
			AssertEqual(t, tc.wantResult, result)
		})
	}
}

type testCloneObj2 struct {
	text string
	Text string
}

func TestClone_WithUnexportedField(t *testing.T) {
	tests := []struct {
		name              string
		toCloneUnexported interface{}
		wantResult        interface{}
	}{
		{"OK", testCloneObj2{text: "123", Text: "567"}, testCloneObj2{text: "123", Text: "567"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Clone(tc.toCloneUnexported)
			AssertEqual(t, tc.wantResult, result, cmp.AllowUnexported(testCloneObj2{}))
		})
	}
}

type testCloneObj3 struct {
	Text *string
}

func TestClone_WithPointerField(t *testing.T) {
	tests := []struct {
		name       string
		toClone    *testCloneObj3
		wantResult *testCloneObj3
	}{
		{"OK, clone pointer struct", &testCloneObj3{Text: Ptr("567")}, &testCloneObj3{Text: Ptr("567")}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Clone(tc.toClone)
			*tc.toClone.Text = "999" // changing original field should not affect cloned
			AssertEqual(t, tc.wantResult, result)
		})
	}
}

func TestPtrString(t *testing.T) {
	str := "123"
	validPtrStr := &str
	tests := []struct {
		name       string
		str        string
		wantResult *string
	}{
		{name: "OK", str: str, wantResult: validPtrStr},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Ptr(tc.str)
			AssertEqual(t, tc.wantResult, result)
		})
	}
}

func TestPtrInt(t *testing.T) {
	num := 123
	validPtrInt := &num
	tests := []struct {
		name       string
		int        int
		wantResult *int
	}{
		{name: "OK", int: num, wantResult: validPtrInt},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Ptr(tc.int)
			AssertEqual(t, tc.wantResult, result)
		})
	}
}

func TestPtrInt64(t *testing.T) {
	num := int64(123)
	validPtrInt64 := &num
	tests := []struct {
		name       string
		int64      int64
		wantResult *int64
	}{
		{name: "OK", int64: num, wantResult: validPtrInt64},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Ptr(tc.int64)
			AssertEqual(t, tc.wantResult, result)
		})
	}
}

func TestPtrFloat(t *testing.T) {
	num := 123.2
	validPtrFloat := &num
	tests := []struct {
		name       string
		float      float64
		wantResult *float64
	}{
		{name: "OK", float: num, wantResult: validPtrFloat},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Ptr(tc.float)
			AssertEqual(t, tc.wantResult, result)
		})
	}
}

func TestPtrFloat32(t *testing.T) {
	num := float32(123.2)
	validPtrFloat := &num
	tests := []struct {
		name       string
		float      float32
		wantResult *float32
	}{
		{name: "OK", float: num, wantResult: validPtrFloat},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Ptr(tc.float)
			AssertEqual(t, tc.wantResult, result)
		})
	}
}
