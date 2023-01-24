// Package listx provides commonly used func for slice comparison
package listx

import (
	"strings"
)

type Equalor[T any] interface {
	Equal(T) bool
}

// Equal return true if both slice a and b are equal
// If elem in slice implemented Equalor interface then they will be compared using Equal method.
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !compare(a[i], b[i]) {
			return false
		}
	}
	return true
}

// IsIntersected return true if any elem existed in s1 also existed in s2.
// If elem in slice implemented Equalor interface then they will be compared using Equal method.
// See TestEq in unit-test for more examples.
func IsIntersected[T comparable](s1, s2 []T) bool {
	for _, a := range s1 {
		for _, b := range s2 {
			if compare(a, b) {
				return true
			}
		}
	}

	return false
}

// FindIntersected return true if any elem existed in s1 also existed in s2.
// Any matching elem will be returned.
func FindIntersected[T comparable](s1, s2 []T) (elementsIntersected []T, found bool) {
	for _, a := range s1 {
		for _, b := range s2 {
			if compare(a, b) {
				elementsIntersected = append(elementsIntersected, b)
			}
		}
	}

	found = len(elementsIntersected) > 0
	return
}

// RemoveDuplicates return slice with duplicated removed.
func RemoveDuplicates[T comparable](elements []T) []T {
	// Use map to record duplicates as we find them.
	encountered := map[any]bool{}
	result := make([]T, 0, len(elements))

	for v := range elements {
		if ok := encountered[elements[v]]; !ok {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

var ComparatorStringCaseInsensitive = func(a, b string) bool {
	return strings.EqualFold(a, b)
}

// Contains return true if slice s contains x.
// If elem in listx implemented Equalor interface then they will be compared using Equal method.
// Simple comparison, use listx.Contains(s, x)
// Eg. For comparing string case-insenstitive, use listx.Contains(s, x, listx.ComparatorStringCaseInsensitive)
func Contains[T comparable](s []T, x T, optionalComparator ...func(T, T) bool) bool {
	for _, elem := range s {
		if len(optionalComparator) > 0 {
			if isAllComparatorOk(elem, x, optionalComparator...) {
				return true
			}
		} else if compare(elem, x) {
			return true
		}
	}
	return false
}

func Filter[T comparable](s []T, filterFunc func(T) bool) []T {
	newList := make([]T, 0, 5)
	for _, elem := range s {
		if filterFunc(elem) {
			newList = append(newList, elem)
		}
	}
	return newList
}

func isAllComparatorOk[T any](a, b T, comparator ...func(T, T) bool) bool {
	for _, c := range comparator {
		if c == nil || !c(a, b) {
			return false
		}
	}
	return true
}

func compare[T comparable](a, b T) bool {
	if x, ok := interface{}(a).(Equalor[T]); ok {
		return x.Equal(b)
	}

	return a == b
}
