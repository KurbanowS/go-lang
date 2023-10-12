package main

import (
	"fmt"

	"strconv"
)

func main() {
	a := "104"
	b := 35
	fmt.Println(strconv.Atoi(a))
	fmt.Println(strconv.Itoa(b + 5))
}
