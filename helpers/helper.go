package helpers

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func Encode(s string) string {
	data := base64.StdEncoding.EncodeToString([]byte(s))
	return string(data)
}

func ValidateToken(validTo time.Time) bool {
	today := time.Now()
	if today.Before(validTo) {
		return true
	}

	return false
}

func CheckPassword(pass string) bool {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return false
	}

	return true
}

func NewHashPassword(pass string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func HashToken(str string) string {
	hashed := sha256.Sum256([]byte(str))
	return fmt.Sprintf("%x", hashed)
}
