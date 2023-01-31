package utils

import (
	"github.com/thanhpk/randstr"
	"strings"
)

func FormatEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

func GenEmailVerificationCode() string {
	code := randstr.String(24)
	return code
}
