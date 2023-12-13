package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "myPassword123"

	// Test hashing the password
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword() error = %v", err)
	}
	if len(hashedPassword) == 0 {
		t.Errorf("HashPassword() returned an empty string")
	}
}

func TestVerifyPassword(t *testing.T) {
	password := "myPassword123"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword() error = %v", err)
	}

	// Test verifying the correct password
	if !VerifyPassword(hashedPassword, password) {
		t.Errorf("VerifyPassword() failed to verify the correct password")
	}

	// Test verifying an incorrect password
	if VerifyPassword(hashedPassword, "wrongPassword") {
		t.Errorf("VerifyPassword() verified an incorrect password")
	}
}
