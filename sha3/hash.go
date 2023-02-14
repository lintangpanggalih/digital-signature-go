package sha3

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

func Hash(plaintext string) string {
	h := sha3.New512()
	h.Write([]byte(plaintext))

	return hex.EncodeToString(h.Sum(nil))
}
