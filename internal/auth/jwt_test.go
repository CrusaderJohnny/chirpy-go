package auth

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestJWTFlow(t *testing.T) {
	secret := "my-super-secret-key"
	userID := uuid.New()
	expiry := time.Hour

	token, err := MakeJWT(userID, secret, expiry)
	if err != nil {
		t.Fatalf("Expected no error on MakeJWT, got %v", err)
	}
	if token == "" {
		t.Fatal("Expected a token string, got empty")
	}

	parsedID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("Expected successful validation, got %v", err)
	}

	if parsedID != userID {
		t.Errorf("Expected UUID %v, got %v", userID, parsedID)
	}
}

func TestValidateJWT_Table(t *testing.T) {
	secret := "correct-secret"
	wrongSecret := "wrong-secret"
	userID := uuid.New()

	validToken, _ := MakeJWT(userID, secret, time.Hour)

	expiredToken, _ := MakeJWT(userID, secret, -time.Hour)

	tests := []struct {
		name        string
		tokenString string
		tokenSecret string
		expectedErr error
	}{
		{
			name:        "Expired Token",
			tokenString: expiredToken,
			tokenSecret: secret,
			expectedErr: jwt.ErrTokenExpired,
		},
		{
			name:        "Wrong Secret Key",
			tokenString: validToken,
			tokenSecret: wrongSecret,
			expectedErr: jwt.ErrSignatureInvalid,
		},
		{
			name:        "Malformed Token String",
			tokenString: "not.a.jwt",
			tokenSecret: secret,
			expectedErr: jwt.ErrTokenMalformed,
		},
		{
			name:        "Invalid Subject UUID",
			tokenString: createTokenWithSubject("not-a-uuid", secret),
			tokenSecret: secret,
			expectedErr: errors.New("invalid UUID length: 10"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ValidateJWT(tt.tokenString, tt.tokenSecret)
			if err == nil {
				t.Errorf("Expected error containing '%s', but got nil", tt.expectedErr)
				return
			}
			if tt.name == "Invalid Subject UUID" {
				if err == nil || !strings.Contains(err.Error(), tt.expectedErr.Error()) {
					t.Errorf("Expected error containing '%s', but got %v", tt.expectedErr, err)
				}
				return
			}
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Expected error '%s', got '%s'", tt.expectedErr, err.Error())
			}
		})
	}
}

func createTokenWithSubject(subject, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject: subject,
	})
	signed, _ := token.SignedString([]byte(secret))
	return signed
}
