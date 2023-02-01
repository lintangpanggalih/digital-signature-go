package rsa

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func Verify(signature string, e, n *big.Int) string {
	var decrypted []string
	var Vs []int64
	// var datas []int64
	var Cs []*big.Int

	signs := strings.Split(signature, ":")
	// fmt.Println(signature)
	// signs, _ := hex.DecodeString(signature)
	// fmt.Println(signs)
	// sign_blocks := strsplit(hex.EncodeToString(signs), 4)
	// fmt.Println(sign_blocks)

	for _, v := range signs {
		V, _ := strconv.ParseInt(v, 16, 64) // hex to dec
		Vs = append(Vs, V)

		C := new(big.Int).Exp(big.NewInt(V), e, n)
		Cs = append(Cs, C)
		hexval := C.Text(16)

		// var data []int64
		// data = append(data, 1)
		// for i := 0; i <= int(e.Int64()); i++ {
		// 	data = append(data, data[i]*V%n.Int64())
		// }
		// datas = append(datas, data[e.Int64()])
		// hexval := strconv.FormatInt(data[e.Int64()], 16)
		block_len := len([]rune(hexval))
		if block_len < 4 {
			hexval = "0" + hexval
		}
		decrypted = append(decrypted, hexval)
	}
	// fmt.Println()
	// fmt.Println(Vs)
	// fmt.Println(Cs)
	// fmt.Println("signature", signature)
	// fmt.Println("decrypted sign", strings.Join(decrypted, ""))
	fmt.Println(decrypted)
	return strings.Join(decrypted, "")
}
