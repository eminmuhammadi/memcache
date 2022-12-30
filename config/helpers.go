package config

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"math/rand"
)

const RAND_INT_MIN = 1000000000
const RAND_INT_MAX = 9999999999
const STRING_LENGTH = 10

func _str() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, STRING_LENGTH)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func hash(message string) string {
	hash := sha512.New()
	hash.Write([]byte(message))

	return hex.EncodeToString(hash.Sum(nil))
}

func RandomString() string {
	msg := fmt.Sprintf("%s.%s.%s", _str(), _str(), _str())
	salt := fmt.Sprintf("%d.", rand.Intn(RAND_INT_MAX-RAND_INT_MIN)+RAND_INT_MIN)

	return hash(salt + msg)
}
