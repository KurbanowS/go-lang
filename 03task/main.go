package main

import (
	"fmt"
	"math/big"
)

func main() {
	var r *float64
	L := 35.0
	Pi := 3.14
	R := L / (2 * Pi)
	r = &R
	fmt.Println(big.NewFloat(Pi*(*r**r)).Text('f', 2))

}
