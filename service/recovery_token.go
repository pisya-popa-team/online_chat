package service

import "crypto/rand"

const numbers = "0123456789"

func RecoveryToken() string {
	bytes := make([]byte, 6)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}

	for i, b := range bytes {
		bytes[i] = numbers[b%byte(len(numbers))]
	}

	return string(bytes)
}