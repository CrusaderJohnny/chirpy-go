package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedToken string
		expectError   bool
	}{
		{
			name: "Valid Bearer Token",
			headers: http.Header{
				"Authorization": []string{"Bearer my-secret-token"},
			},
			expectedToken: "my-secret-token",
			expectError:   false,
		},
		{
			name: "Valid Token with extra whitespace",
			headers: http.Header{
				"Authorization": []string{"  Bearer   my-secret-token  "},
			},
			expectedToken: "my-secret-token",
			expectError:   false,
		},
		{
			name:          "Missing Authorization header",
			headers:       http.Header{},
			expectedToken: "",
			expectError:   true,
		},
		{
			name: "Empty Authorization header",
			headers: http.Header{
				"Authorization": []string{""},
			},
			expectedToken: "",
			expectError:   true,
		},
		{
			name: "Non-Bearer Authorization (e.g., Basic)",
			headers: http.Header{
				"Authorization": []string{"Basic dXNlcjpwYXNz"},
			},
			expectedToken: "Basic dXNlcjpwYXNz", // Note: current logic just returns the string
			expectError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GetBearerToken(tt.headers)

			if (err != nil) != tt.expectError {
				t.Errorf("GetBearerToken() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if token != tt.expectedToken {
				t.Errorf("GetBearerToken() = %v, want %v", token, tt.expectedToken)
			}
		})
	}
}
