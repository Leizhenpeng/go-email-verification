package utils

import (
	"encoding/base64"
	"errors"
	"github.com/thanhpk/randstr"
	"strings"
)

func FormatEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

func GenCode() string {
	code := randstr.String(24)
	return code
}
func GenEmailVerificationInfo(email string) string {
	code := GenCode()
	info := Encode(email + "|" + code)
	return info
}

func ParseEmailVerificationInfo(info string) (email string, code string, err error) {
	data, err := Decode(info)
	if err != nil {
		return "", "", err
	}
	s := strings.Split(data, "|")
	if len(s) != 2 {
		return "", "", errors.New("invalid email verification info")
	}
	return s[0], s[1], nil
}

func Encode(s string) string {
	data := base64.StdEncoding.EncodeToString([]byte(s))
	return string(data)
}

func Decode(s string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
