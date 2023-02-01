package prime

import (
	crypto_rand "crypto/rand"
	"fmt"
	"math/big"
)

var zero = big.NewInt(0)
var one = big.NewInt(1)
var two = big.NewInt(2)

func GeneratePrimes(bits int) (*big.Int, *big.Int) {
	p, _ := crypto_rand.Prime(crypto_rand.Reader, bits)
	var q *big.Int
	for {
		q, _ = crypto_rand.Prime(crypto_rand.Reader, bits)
		if q.Cmp(p) < 0 {
			break
		}
	}
	return p, q
}

func CreateKey(p *big.Int, q *big.Int) (*big.Int, *big.Int, *big.Int) {
	e := big.NewInt(0)
	d := big.NewInt(0)
	n := new(big.Int).Mul(p, q)
	m := new(big.Int).Mul(p.Sub(p, one), q.Sub(q, one))

	for i := new(big.Int).Set(two); i.Cmp(m) == -1; i.Add(i, one) {
		gcd := new(big.Int).GCD(nil, nil, i, m)
		if gcd.Cmp(one) == 0 {
			e = i
			break
		}
	}

	// for j := new(big.Int).Set(two); j.Cmp(m) == -1; j.Add(j, one) {
	// 	exp := new(big.Int).Mul(j, e)
	// 	if new(big.Int).Mod(exp, m).Cmp(one) == 0 {
	// 		d = j
	// 		break
	// 	}
	// }

	var x, y big.Int
	y.GCD(&x, nil, e, m)

	// if x is negative
	if x.Cmp(zero) < 0 {
		d.Add(&x, m)
	} else {
		d.Set(&x)
	}
	fmt.Println("m", m)
	fmt.Println("n", n)
	fmt.Println("e", e)
	fmt.Println("d", d)

	return n, e, d
}
