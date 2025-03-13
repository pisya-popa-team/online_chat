package utils

import (
	"strconv"
	"time"
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

var layout = "00.00.00 00:00"
func FormatStringToDate(s string) time.Time{
	t, _ := time.Parse(layout, s)
	return t
}