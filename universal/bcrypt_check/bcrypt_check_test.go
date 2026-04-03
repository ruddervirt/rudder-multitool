package bcryptcheck

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestCheckBcrypt_Success(t *testing.T) {
	password := []byte("correcthorsebatterystaple")
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to generate hash: %v", err)
	}

	if err := CheckBcrypt(string(hash), string(password)); err != nil {
		t.Fatalf("expected password to match, got error: %v", err)
	}
}

func TestCheckBcrypt_Failure(t *testing.T) {
	password := []byte("password1")
	wrong := []byte("password2")
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		t.Fatalf("failed to generate hash: %v", err)
	}

	if err := CheckBcrypt(string(hash), string(wrong)); err == nil {
		t.Fatalf("expected mismatch error, got nil")
	}
}
