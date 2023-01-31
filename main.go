package main

import (
	"digital-signature/rsa"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	signature := rsa.Sign("Hello World")
	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s\n", elapsed)

	fmt.Println(signature)
}
