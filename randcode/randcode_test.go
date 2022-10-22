package randcode

import (
	"testing"

	"github.com/urpent/go/ut"
)

func TestSecure(t *testing.T) {
	tests := []struct {
		name       string
		codeLen    int
		wantResult int
		wantErr    error
	}{
		{
			name:       "Ok",
			codeLen:    15,
			wantResult: 15,
			wantErr:    nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Secure(tc.codeLen)
			ut.AssertEqual(t, tc.wantResult, len(result))
			ut.AssertEqual(t, tc.wantErr, err)
		})
	}
}

func TestCustomSecure(t *testing.T) {
	tests := []struct {
		name        string
		codeLen     int
		combination Combination
		wantResult  int
		wantErr     error
	}{
		{
			name:        "Ok",
			codeLen:     15,
			combination: "abc",
			wantResult:  15,
			wantErr:     nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := CustomSecure(tc.codeLen, tc.combination)
			ut.AssertEqual(t, tc.wantResult, len(result))
			ut.AssertEqual(t, tc.wantErr, err)
		})
	}
}
