package utils

import "strings"

func ExtractTokenFromHeaderString(header string) string {
	header_parts := strings.Split(header, " ")
	return header_parts[1]
}