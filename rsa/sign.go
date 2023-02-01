package rsa

import (
	"bytes"
	"digital-signature-go/prime"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"golang.org/x/crypto/sha3"
)

func Sign(plaintext string) (string, string, *big.Int, *big.Int) {
	var signs []string
	var Ms []*big.Int
	var Vs []int64
	// var datas []int64

	p, q := prime.GeneratePrimes(16)
	n, e, d := prime.CreateKey(p, q)

	h := sha3.New512()
	h.Write([]byte(plaintext))
	hashed := h.Sum(nil)

	chiper := hex.EncodeToString(hashed)
	chiper_blocks := strsplit(chiper, 4)

	for _, v := range chiper_blocks {
		V, _ := strconv.ParseInt(v, 16, 64)
		Vs = append(Vs, V)

		M := new(big.Int).Exp(big.NewInt(V), d, n)
		Ms = append(Ms, M)
		hexval := M.Text(16)

		// var data []int64
		// data = append(data, 1)
		// for i := 0; i <= int(d.Int64()); i++ {
		// 	data = append(data, data[i]*V%n.Int64())
		// }
		// datas = append(datas, data[d.Int64()])
		// hexval := strconv.FormatInt(data[d.Int64()], 16)

		signs = append(signs, hexval)
	}
	// fmt.Println("original chiper", chiper)
	// fmt.Println(Vs)
	// fmt.Println(Ms)
	fmt.Println(chiper_blocks)
	// fmt.Println(hex.DecodeString(strings.Join(signs, "")))
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
