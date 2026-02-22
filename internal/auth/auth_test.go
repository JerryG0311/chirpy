package auth

import (
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
