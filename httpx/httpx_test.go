package httpx

import (
	"net/http"
	"testing"

	"github.com/urpent/go/ut"
)

func Test_httpClient_buildRequestHeader(t *testing.T) {
	tests := []struct {
		name       string
		client     httpClient
		header     http.Header
		wantResult http.Header
	}{
		{
			name: "Ok, array is equal",
			client: httpClient{defaultHeader: map[string][]string{
				"test": {"a", "z"},
			}},
			header: map[string][]string{
				"test": {"a", "z1"},
			},
			wantResult: map[string][]string{
				"Test": {"a", "z", "a", "z1"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.client.buildRequestHeader(tc.header)
			ut.AssertEqual(t, tc.wantResult, result)
		})
	}
}
