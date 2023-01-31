package rsa

import (
	"bytes"
	"digital-signature-go/prime"
	"encoding/hex"
	"math/big"
	"strconv"
	"strings"

	"golang.org/x/crypto/sha3"
)

func Sign(plaintext string) (string, *big.Int, *big.Int) {
	var signs []string
	h := sha3.New512()

	h.Write([]byte(plaintext))
	hashed := h.Sum(nil)

	chiper := hex.EncodeToString(hashed)
	chiper_blocks := strsplit(chiper, 4)

	p, q := prime.GeneratePrimes(8)
	n, e, d := prime.CreateKey(p, q)
	for _, v := range chiper_blocks {
		V, _ := strconv.ParseInt(v, 16, 64)
		M := new(big.Int).Exp(big.NewInt(V), d, n)
		hexval := M.Text(16)
		signs = append(signs, hexval)
	}
	return strings.Join(signs, ":"), n, e
}
func strsplit(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}
