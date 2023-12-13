package utils

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestSignToken(t *testing.T) {
	jwtSecret := "mysecret"
	sub := "testuser"
	tokenString, err := SignToken(sub, jwtSecret)
	if err != nil {
		t.Fatalf("Failed to sign token: %v", err)
	}

	if tokenString == "" {
		t.Fatalf("Token string is empty")
	}
}

func TestVerifyToken(t *testing.T) {
	jwtSecret := "mysecret"
	sub := "testuser"
	tokenString, _ := SignToken(sub, jwtSecret)

	// Test valid token
	token, err := VerifyToken(tokenString, jwtSecret)
	if err != nil {
		t.Fatalf("Failed to verify token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["sub"] != sub {
			t.Errorf("Expected sub %v, got %v", sub, claims["sub"])
		}
	} else {
		t.Fatalf("Token claims are not valid")
	}

	// Test invalid token
	_, err = VerifyToken("invalid.token.string", jwtSecret)
	if err == nil {
		t.Fatalf("Expected error verifying invalid token, got none")
	}
}
