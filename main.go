package main

import (
	"digital-signature-go/rsa"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	digest, signature, e, n := rsa.Sign("Sasenna Lintang")
	decrypted := rsa.Verify(signature, e, n)
	fmt.Println(digest)
	fmt.Println(decrypted)
	if digest == decrypted {
		fmt.Println("Valid")
	} else {
		fmt.Println("Not Valid")
	}
	// fmt.Println(new(big.Int).Exp(big.NewInt(8372), big.NewInt(11), big.NewInt(50429)))
	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s\n", elapsed)
}
