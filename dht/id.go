package dht

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(n int8) string {
	buf := make([]byte, n)
	for i := range buf {
		buf[i] = letters[rand.Intn(len(letters))]
	}
	return string(buf)
}

func randomID() []byte {
	newSha := sha1.New()
	newSha.Write([]byte(randomString(20)))
	return newSha.Sum(nil)
}

func idToString(id []byte) string {
	return hex.EncodeToString(id)
}
