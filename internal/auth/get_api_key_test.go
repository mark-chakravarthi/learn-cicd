package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		header  string
		want    string
		wantErr bool
		errIs   error
	}{
		"valid header": {header: "ApiKey abc123", want: "abc123"},
		"no header":    {header: "", wantErr: true, errIs: ErrNoAuthHeaderIncluded},
		"wrong scheme": {header: "Bearer abc123", wantErr: false},
		"missing key":  {header: "ApiKey", wantErr: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			headers := http.Header{}
			if tc.header != "" {
				headers.Set("Authorization", tc.header)
			}

			got, err := GetAPIKey(headers)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got key %q", got)
				}
				if tc.errIs != nil && !errors.Is(err, tc.errIs) {
					t.Fatalf("expected error %v, got %v", tc.errIs, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("expected %q, got %q", tc.want, got)
			}
		})
	}
}
