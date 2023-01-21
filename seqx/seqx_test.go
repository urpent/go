package seqx

import (
	"testing"

	"github.com/urpent/go/ut"
)

func TestMiddle(t *testing.T) {
	tests := []struct {
		name       string
		prev       string
		next       string
		wantResult string
	}{
		{
			name:       "OK",
			prev:       "n",
			next:       "",
			wantResult: "u",
		},
		{
			name:       "OK",
			prev:       "z",
			next:       "",
			wantResult: "zn",
		},
		{
			name:       "OK",
			prev:       "a",
			next:       "",
			wantResult: "n",
		},
		{
			name:       "OK",
			prev:       "z",
			next:       "z",
			wantResult: "zn",
		},
		{
			name:       "OK",
			prev:       "a",
			next:       "ab",
			wantResult: "aan",
		},
		{
			name:       "OK",
			prev:       "a",
			next:       "aab",
			wantResult: "aaan",
		},
		{
			name:       "OK",
			prev:       "a",
			next:       "b",
			wantResult: "an",
		},
		{
			name:       "OK",
			prev:       "a",
			next:       "an",
			wantResult: "ag",
		},
		{
			name:       "OK",
			prev:       "abcde",
			next:       "abchi",
			wantResult: "abcf",
		},
		{
			name:       "OK",
			prev:       "abc",
			next:       "abchi",
			wantResult: "abcd",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Middle(tc.prev, tc.next)
			ut.AssertEqual(t, tc.wantResult, result)
		})
	}
}

func TestMiddle_Unsupported_case(t *testing.T) {
	tests := []struct {
		name       string
		prev       string
		next       string
		wantResult string
	}{
		{
			name:       "Return wrong string, prev & next should never be equal",
			prev:       "ab",
			next:       "ab",
			wantResult: "abn", // sort is wrong -> ab, ab, abn
		},
		{
			name:       "Return wrong string, string should never end with a",
			prev:       "a",
			next:       "aa",
			wantResult: "aan", // sort is wrong -> a, aa, aan
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Middle(tc.prev, tc.next)
			ut.AssertEqual(t, tc.wantResult, result)
		})
	}
}

func TestMiddle_Consecutive_characters(t *testing.T) {
	tests := []struct {
		name       string
		prev       string
		next       string
		wantResult string
	}{
		{
			name:       "OK",
			prev:       "abhs",
			next:       "abit",
			wantResult: "abhw",
		},
		{
			name:       "OK",
			prev:       "abh",
			next:       "abit",
			wantResult: "abhn",
		},
		{
			name:       "OK",
			prev:       "abhz",
			next:       "abit",
			wantResult: "abhzn",
		},
		{
			name:       "OK",
			prev:       "abhzs",
			next:       "abit",
			wantResult: "abhzw",
		},
		{
			name:       "OK",
			prev:       "abhzz",
			next:       "abit",
			wantResult: "abhzzn",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Middle(tc.prev, tc.next)
			ut.AssertEqual(t, tc.wantResult, result)
		})
	}
}

func TestMiddle_Right_character_is_a_or_b(t *testing.T) {
	tests := []struct {
		name       string
		prev       string
		next       string
		wantResult string
	}{
		{
			name:       "OK",
			prev:       "abc",
			next:       "abcah",
			wantResult: "abcad",
		},
		{
			name:       "OK",
			prev:       "abc",
			next:       "abcab",
			wantResult: "abcaan",
		},
		{
			name:       "OK",
			prev:       "abc",
			next:       "abcaah",
			wantResult: "abcaad",
		},
		{
			name:       "OK",
			prev:       "abc",
			next:       "abcb",
			wantResult: "abcan",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Middle(tc.prev, tc.next)
			ut.AssertEqual(t, tc.wantResult, result)
		})
	}
}
