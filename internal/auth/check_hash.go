package auth

import "github.com/alexedwards/argon2id"

func CheckHashPassword(password string, hash string) (bool, error) {
	found, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, err
	}
	return found, nil
}
