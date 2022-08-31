package main

import (
	crand "crypto/rand"
	"math/big"
	mrand "math/rand"
	"time"
)

func main() {
	mrand.Seed(time.Now().UnixNano())
	mrand.Intn(5)
	mrand.Float32()

	crand.Int(crand.Reader, big.NewInt(10000))
}
