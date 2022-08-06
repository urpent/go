package listx

import (
	"strings"
	"testing"

	"github.com/urpent/go/ut"
)

type TestEq struct {
	ID int
	X  int
}

func (t TestEq) Equal(t2 TestEq) bool {
	return t.ID == t2.ID
}

func TestEqual_Equalor(t *testing.T) {
	tests := []struct {
		name       string
		arrA       []TestEq
		arrB       []TestEq
		wantResult bool
	}{
		{
			name:       "Ok, array is equal",
			arrA:       []TestEq{{ID: 2, X: 1}, {ID: 1, X: 2}},
			arrB:       []TestEq{{ID: 2, X: 3}, {ID: 1, X: 4}},
			wantResult: true,
		},
		{
			name:       "Not equal, different length",
			arrA:       []TestEq{{ID: 2}, {ID: 1}},
			arrB:       []TestEq{{ID: 1}},
			wantResult: false,
		},
		{
			name:       "Not equal, different order",
			arrA:       []TestEq{{ID: 2}, {ID: 1}},
			arrB:       []TestEq{{ID: 1}, {ID: 2}},
			wantResult: false,
		},
		{
			name:       "Not equal, wrong id",
			arrA:       []TestEq{{ID: 2}, {ID: 1}},
			arrB:       []TestEq{{ID: 1}, {ID: 1}},
			wantResult: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Equal(tc.arrA, tc.arrB)
			ut.AssertEqual(t, tc.wantResult, result)
		})
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		name       string
		arrA       []int
		arrB       []int
		wantResult bool
	}{
		{
			name:       "Ok, array is equal",
			arrA:       []int{2, 1},
			arrB:       []int{2, 1},
			wantResult: true,
		},
		{
			name:       "Not equal, different length",
			arrA:       []int{2, 1},
			arrB:       []int{2},
			wantResult: false,
		},
		{
			name:       "Not equal, different order",
			arrA:       []int{2, 1},
			arrB:       []int{1, 2},
			wantResult: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Equal(tc.arrA, tc.arrB)
			ut.AssertEqual(t, tc.wantResult, result)
		})
	}
}

func TestIsIntersected_Equalor(t *testing.T) {
	tests := []struct {
		name       string
		arrA       []TestEq
		arrB       []TestEq
		wantResult bool
	}{
		{
			name:       "Ok",
			arrA:       []TestEq{{ID: 1}, {ID: 3, X: 1}},
			arrB:       []TestEq{{ID: 2}, {ID: 3, X: 9}},
			wantResult: true,
		},
		{
			name:       "Not intersected",
			arrA:       []TestEq{{ID: 1}, {ID: 2}},
			arrB:       []TestEq{{ID: 4}, {ID: 5}},
			wantResult: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := IsIntersected(tc.arrA, tc.arrB)
			ut.AssertEqual(t, tc.wantResult, result)
		})
	}
}

func TestIsIntersected_Int64(t *testing.T) {
	tests := []struct {
		name       string
		arrA       []int64
		arrB       []int64
		wantResult bool
	}{
		{name: "Ok", arrA: []int64{1, 3}, arrB: []int64{2, 3}, wantResult: true},
		{name: "Not intersected", arrA: []int64{1, 2}, arrB: []int64{4, 5}, wantResult: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := IsIntersected(tc.arrA, tc.arrB)
			ut.AssertEqual(t, tc.wantResult, result)
		})
	}
}

func TestFindIntersected_String(t *testing.T) {
	tests := []struct {
		name            string
		arrA            []string
		arrB            []string
		wantResult      []string
		wantIntersected bool
	}{
		{
			name:            "Ok",
			arrA:            []string{"1", "3", "4"},
			arrB:            []string{"2", "3", "4"},
			wantResult:      []string{"3", "4"},
			wantIntersected: true,
		},
		{
			name:            "Not intersected",
			arrA:            []string{"1", "2"},
			arrB:            []string{"4", "5"},
			wantIntersected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, isIntersected := FindIntersected(tc.arrA, tc.arrB)
			ut.AssertEqual(t, tc.wantResult, result)
			ut.AssertEqual(t, tc.wantIntersected, isIntersected)
		})
	}
}

func TestRemoveDuplicates_Int64(t *testing.T) {
	tests := []struct {
		name       string
		array      []int64
		wantResult []int64
	}{
		{
			name:       "Ok, remove duplicate",
			array:      []int64{5, 1, 2, 5, 3, 3, 4},
			wantResult: []int64{5, 1, 2, 3, 4},
		},
		{
			name:       "Ok, no duplicate",
			array:      []int64{3, 4, 1},
			wantResult: []int64{3, 4, 1},
		},
		{
			name:       "Ok, empty array",
			array:      []int64{},
			wantResult: []int64{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := RemoveDuplicates(tc.array)
			ut.AssertEqual(t, tc.wantResult, result)
		})
	}
}

func TestRemoveDuplicates_String(t *testing.T) {
	tests := []struct {
		name       string
		array      []string
		wantResult []string
	}{
		{
			name:       "Ok, remove duplicate",
			array:      []string{"5", "1", "2", "5", "3", "3", "4"},
			wantResult: []string{"5", "1", "2", "3", "4"},
		},
		{
			name:       "Ok, no duplicate",
			array:      []string{"3", "4", "1"},
			wantResult: []string{"3", "4", "1"},
		},
		{
			name:       "Ok, empty array",
			array:      []string{},
			wantResult: []string{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := RemoveDuplicates(tc.array)
			ut.AssertEqual(t, tc.wantResult, result)
		})
	}
}

func TestContains_Int64(t *testing.T) {
	tests := []struct {
		name       string
		array      []int64
		i          int64
		wantResult bool
	}{
		{
			name:       "Ok",
			array:      []int64{1, 2},
			i:          2,
			wantResult: true,
		},
		{
			name:  "Not Ok",
			array: []int64{1, 2},
			i:     3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Contains(tc.array, tc.i)
			ut.AssertEqual(t, tc.wantResult, result)
		})
	}
}

func TestContains_String(t *testing.T) {
	tests := []struct {
		name       string
		array      []string
		s          string
		wantResult bool
	}{
		{
			name:       "Ok",
			array:      []string{"1", "2"},
			s:          "2",
			wantResult: true,
		},
		{
			name:  "Not Ok",
			array: []string{"1", "2"},
			s:     "3",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Contains(tc.array, tc.s)
			ut.AssertEqual(t, tc.wantResult, result)
		})
	}
}

func TestContains_OptionalComparatorStringCaseInsensitive(t *testing.T) {
	tests := []struct {
		name             string
		array            []string
		s                string
		customComparator func(string, string) bool
		wantResult       bool
	}{
		{
			name:             "Ok, case insensitive",
			array:            []string{"ABC", "efg"},
			s:                "abc",
			customComparator: ComparatorStringCaseInsensitive,
			wantResult:       true,
		},
		{
			name:             "Not Ok",
			array:            []string{"abc", "efg"},
			s:                "abc",
			customComparator: nil,
			wantResult:       false,
		},
		{
			name:             "Not Ok",
			array:            []string{"ABC", "efg"},
			s:                "abc",
			customComparator: nil,
			wantResult:       false,
		},
		{
			name:             "Not Ok",
			array:            []string{"1", "2"},
			s:                "3",
			customComparator: nil,
			wantResult:       false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Contains(tc.array, tc.s, tc.customComparator)
			ut.AssertEqual(t, tc.wantResult, result)
		})
	}
}

func TestContains_MultipleComparator(t *testing.T) {
	tests := []struct {
		name              string
		array             []string
		s                 string
		customComparator1 func(string, string) bool
		customComparator2 func(string, string) bool
		wantResult        bool
	}{
		{
			name:  "Ok, elem contains substring with length more than 4",
			array: []string{"abc", "abcdefg", "efg"},
			s:     "ab",
			customComparator1: func(a, b string) bool {
				return strings.Contains(a, b)
			},
			customComparator2: func(a, b string) bool {
				return len(a) > 4
			},
			wantResult: true,
		},
		{
			name:  "Not Ok, one comparison failed, contains substring but length but more than 4",
			array: []string{"abc", "efg"},
			s:     "ab",
			customComparator1: func(a, b string) bool {
				return strings.Contains(a, b)
			},
			customComparator2: func(a, b string) bool {
				return len(a) > 4
			},
			wantResult: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Contains(tc.array, tc.s, tc.customComparator1, tc.customComparator2)
			ut.AssertEqual(t, tc.wantResult, result)
		})
	}
}
