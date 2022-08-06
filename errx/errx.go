package errx

import (
	"errors"
	"strings"
)

const DefaultSep = ":\n"

// Wrap returns the result of wrapped array of errors.
// First error will be the head
func Wrap(errs ...error) error {
	return chain(errs...)
}

// chain returns the result of wrapped array of errors.
// First error will be the head
func chain(errs ...error) error {
	errCount := len(errs)

	if errCount == 0 {
		return nil
	} else if errCount == 1 {
		return errs[0]
	}

	head := &wrapErr{err: errs[0]}
	current := head
	for i := 1; i < errCount; i++ {
		next := &wrapErr{err: errs[i]}
		current.innerErr = next
		current = next
	}

	return head
}

type wrapErr struct {
	err      error
	innerErr error
}

func (e wrapErr) Error() string {
	if e.err == nil {
		return ""
	}

	errs := []string{e.err.Error()}
	if e.innerErr != nil {
		errs = append(errs, e.innerErr.Error())
	}

	return strings.Join(errs, DefaultSep)
}

func (e wrapErr) Is(target error) bool {
	return errors.Is(target, e.err)
}

func (e wrapErr) Unwrap() error {
	return e.innerErr
}
