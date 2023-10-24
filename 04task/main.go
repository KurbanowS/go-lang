package main

import "fmt"

func main() {
	weekend := []int{1, 2, 3, 4, 5, 6, 7}
	workdays := weekend[:5]
	weekend = weekend[5:7]
	fmt.Println(workdays, weekend)
	week := append(workdays, weekend...)
	fmt.Println(week)
}
