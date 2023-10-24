package main

import "fmt"

func linearsearch(a []string, x string) bool {
	for _, item := range a {
		if item == x {
			fmt.Println("We have this word")
			return true
		}
	}
	fmt.Println("We don't have this word")
	return false
}

func getMax(numbs []int) int {
	for i := range numbs {
		if numbs[0] < numbs[i] {
			numbs[0] = numbs[i]
		}
	}
	fmt.Println(numbs[0])
	return numbs[0]
}

func main() {
	items := []string{"hello", "world", "house", "yeyo", "bool", "rent"}
	fmt.Println(linearsearch(items, "bool"))
	getMax([]int{1, 2, 3, 6, 14, 46})

}
