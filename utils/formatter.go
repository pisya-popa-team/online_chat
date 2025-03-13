package utils

import (
	"strconv"
)

func IntToString(u int) string {
	s := strconv.Itoa(u)
    return s
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
    return i
}

func PointerTo[T ~string](s T) *T {
    return &s
}