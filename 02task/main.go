package main

import (
	"fmt"
	"math/big"
)

func main() {
	type EuropeanVelocity float64
	type AmericanVelocity float64
	var toeur EuropeanVelocity = 3.6
	var toamer AmericanVelocity = 2.24
	a := 120.4
	esum := a * float64(toeur)
	b := 130.0
	asum := b * float64(toamer)
	fmt.Println(big.NewFloat(esum).Text('f', 2))
	fmt.Println(big.NewFloat(asum).Text('f', 2))
}
