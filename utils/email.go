package utils

import "strings"

func FormatEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}
