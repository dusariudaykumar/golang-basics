package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Go lang math func")
	numberOne := 1
	numberTwo := 2.20
	fmt.Println(float64(numberOne) + numberTwo)

	//random num from rand
	// fmt.Println("randomNum: ", rand.Intn(10))

	//random num from crypto
	cryptoRandNum, _ := rand.Int(rand.Reader, big.NewInt(10))
	fmt.Println(cryptoRandNum)
}
