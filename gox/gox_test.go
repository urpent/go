package gox

import (
	"testing"

	"github.com/urpent/go/ut"
)

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
			ut.AssertEqual(t, tc.wantResult, result)
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
			ut.AssertEqual(t, tc.wantResult, result)
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
			ut.AssertEqual(t, tc.wantResult, result)
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
			ut.AssertEqual(t, tc.wantResult, result)
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
			ut.AssertEqual(t, tc.wantResult, result)
		})
	}
}
