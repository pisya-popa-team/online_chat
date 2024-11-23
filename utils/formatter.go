package utils

import "strconv"

func StringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}

func UintToString(u uint) string {
	s := strconv.Itoa(int(u))
	return s
}

func PointerTo[T ~string](s T) *T {
    return &s
}