package errx

import (
	"errors"
	"testing"

	"github.com/urpent/go/ut"
)

func TestErrWrap_errors_Is(t *testing.T) {
	e1 := errors.New("e1")
	e2 := errors.New("e2")
	e3 := errors.New("e3")
	e4 := errors.New("e4")

	tests := []struct {
		name       string
		errs       []error
		containErr error
		wantResult bool
	}{
		{
			name:       "Ok, contain e1",
			errs:       []error{e1, e2, e3},
			containErr: e1,
			wantResult: true,
		},
		{
			name:       "Ok, contain e2",
			errs:       []error{e1, e2, e3},
			containErr: e2,
			wantResult: true,
		},
		{
			name:       "Ok, contain e3",
			errs:       []error{e1, e2, e3},
			containErr: e3,
			wantResult: true,
		},
		{
			name:       "Does not contain error",
			errs:       []error{e1, e2, e3},
			containErr: e4,
			wantResult: false,
		},
		{
			name:       "Does not contain error",
			errs:       []error{nil, nil, nil},
			containErr: e4,
			wantResult: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			wrappedError := Wrap(tc.errs...)
			result := errors.Is(wrappedError, tc.containErr)
			ut.AssertEqual(t, tc.wantResult, result)
		})
	}

	t.Run("Ok, wrap nil error return nil", func(t *testing.T) {
		result := Wrap(nil)
		ut.AssertEqual(t, nil, result)
	})

	t.Run("Ok, calling wrap with nothing return nil", func(t *testing.T) {
		result := Wrap()
		ut.AssertEqual(t, nil, result)
	})
}

func TestErrWrap_Error(t *testing.T) {
	e1 := errors.New("e1")
	e2 := errors.New("e2")
	e3 := errors.New("e3")

	tests := []struct {
		name       string
		errs       []error
		wantResult string
	}{
		{
			name:       "Print all errors",
			errs:       []error{e1, e2, e3},
			wantResult: "e1:\ne2:\ne3",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			wrappedError := Wrap(tc.errs...)
			result := wrappedError.Error()
			ut.AssertEqual(t, tc.wantResult, result)
		})
	}
}
