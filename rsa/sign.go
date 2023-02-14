package rsa

import (
	"bytes"
	"digital-signature-go/prime"
	"digital-signature-go/sha3"
	"math/big"
	"strconv"
	"strings"
)

func Sign(plaintext string) (string, string, *big.Int, *big.Int) {
	var signs []string

	p, q := prime.GeneratePrimes(16)
	n, e, d := prime.CreateKey(p, q)

	chiper := sha3.Hash(plaintext)

	chiper_blocks := strsplit(chiper, 4)

	for _, v := range chiper_blocks {
		V, _ := strconv.ParseInt(v, 16, 64)

		M := new(big.Int).Exp(big.NewInt(V), d, n)
		hexval := M.Text(16)

		signs = append(signs, hexval)
	}
	return chiper, strings.Join(signs, ":"), e, n
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
