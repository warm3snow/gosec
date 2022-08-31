package main

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	mrand "math/rand"
	"time"
)

func rand() {
	mrand.Seed(time.Now().UnixNano())
	mrand.Intn(5)
	mrand.Float32()

	crand.Int(crand.Reader, big.NewInt(10000))
}

func goroutine() {
	go func() {
		fmt.Println("this is goroutine")
	}()
}

func timeTest() {
	time.Sleep(time.Second)
	time.Now()
	time.Unix(1, 1)
}
