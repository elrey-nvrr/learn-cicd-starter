package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		input http.Header
		want  string
		err   error
	}{
		{
			input: http.Header{"Authorization": {"ApiKey 12345"}},
			want:  "12345",
			err:   nil,
		},
		{
			input: http.Header{"Authorization": {"Bearer 12345"}},
			want:  "",
			err:   errors.New("malformed authorization header"),
		},
		{
			input: http.Header{"Auth": {"ApiKey abcdef"}},
			want:  "",
			err:   errors.New("no authorization header included"),
		},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if got != tc.want || (err != nil && err.Error() != tc.err.Error()) {
			t.Errorf("GetAPIKey(%v) = %v, %v; want %v, %v", tc.input, got, err, tc.want, tc.err)
		}
	}
}
