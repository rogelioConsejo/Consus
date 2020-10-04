package persistence

import (
	"encoding/hex"
	"math/rand"
	"time"
)

func CreateToken(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	if _, err := r.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
