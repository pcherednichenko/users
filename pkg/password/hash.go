package password

import "golang.org/x/crypto/bcrypt"

// Hash returns from password using bcrypt library
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
