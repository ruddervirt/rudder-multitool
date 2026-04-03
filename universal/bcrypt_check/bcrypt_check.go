package bcryptcheck

import "golang.org/x/crypto/bcrypt"

// CheckBcrypt compares a bcrypt hash and plaintext password.
// Returns nil if match, otherwise an error.
func CheckBcrypt(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
