package rsa

import (
	"math/big"
	"strconv"
	"strings"
)

func Verify(signature string, e, n *big.Int) string {
	var decrypted []string

	signs := strings.Split(signature, ":")

	for _, v := range signs {
		V, _ := strconv.ParseInt(v, 16, 64) // hex to dec
		C := new(big.Int).Exp(big.NewInt(V), e, n)
		hexval := C.Text(16)

		block_len := len([]rune(hexval))
		if block_len < 4 {
			hexval = "0" + hexval
		}
		decrypted = append(decrypted, hexval)
	}
	return strings.Join(decrypted, "")
}
