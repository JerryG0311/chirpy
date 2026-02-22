package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestJWT(t *testing.T) {
	secret := "secret"
	userID := uuid.New()
	expirationTime := time.Hour

	token, err := MakeJWT(userID, secret, expirationTime)
	if err != nil {
		t.Fatalf("unexpected error making JWT: %v", err)
	}

	returnedUserID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("unexpected error validating JWT: %v", err)
	}

	if returnedUserID != userID {
		t.Errorf("expected userID %v, got %v", userID, returnedUserID)
	}
}

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedToken string
		expectedErr   bool
	}{
		{
			name: "Valid Bearer Token",
			headers: http.Header{
				"Authorization": []string{"Bearer my-secret-token"},
			},
			expectedToken: "my-secret-token",
			expectedErr:   false,
		},
		{
			name:          "Missing Authorization Header",
			headers:       http.Header{},
			expectedToken: "",
			expectedErr:   true,
		},
		{
			name: "Malformed Header (Missing Bearer)",
			headers: http.Header{
				"Authorization": []string{"NotBearer some-token"},
			},
			expectedToken: "",
			expectedErr:   true,
		},
		{
			name: "Malformed Header (Missing Token)",
			headers: http.Header{
				"Authorization": []string{"Bearer"},
			},
			expectedToken: "",
			expectedErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GetBearerToken(tt.headers)
			if (err != nil) != tt.expectedErr {
				t.Errorf("GetBearerToken() error = %v, expectedErr %v", err, tt.expectedErr)
				return
			}
			if token != tt.expectedToken {
				t.Errorf("GetBearerToken() = %v, expected %v", token, tt.expectedToken)
			}
		})
	}
}
