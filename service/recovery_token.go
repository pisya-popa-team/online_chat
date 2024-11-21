package service

import "math/rand"

func RecoveryToken() int {
	min := 100000
	max := 999999
	return rand.Intn(max-min) + min
}