package utils

import (
	mathRand "math/rand"
	"time"
)

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type HexId struct {
	Id string
}

type HexAction interface {
	RandomString(length int)
	SecureRandomString(length int)
}

func (idHex *HexId) RandomString(length int) {
	result := make([]rune, length)
	mathRand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		result[i] = alphabet[mathRand.Intn(len(alphabet))]
	}
	idHex.Id = string(result[:length])
}
