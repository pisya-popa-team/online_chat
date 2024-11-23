package utils

import (
	"strconv"
)

func IntToString(u int) string {
	s := strconv.Itoa(u)
    return s
}

func PointerTo[T ~string](s T) *T {
    return &s
}