package service

import "crypto/rand"

const alphabet = "0123456789"

func RecoveryToken() string {
	bytes := make([]byte, 6)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}

	for i, b := range bytes {
		bytes[i] = alphabet[b%byte(len(alphabet))]
	}

	return string(bytes)
}