package main

import "fmt"

// func words(w ...string) []string {
// 	fmt.Println(printWord(w...))
// 	return w
// }

// func printWord(w ...string) int {
// 	for num, word := range w {
// 		fmt.Println(num, word)
// 	}
// 	return len(w)
// }
type wordLogger func(w ...string) (cnt int)

func main() {

	printWord := func(w ...string) (cnt int) {
		for _, word := range w {
			fmt.Println(word)
		}

		return len(w)
	}

	func(printer wordLogger, w ...string) {
		fmt.Printf("Count words: %d\n", printer(w...))
	}(printWord, "str1", "str2")
	// author := map[string]map[string]string{
	// 	"John": {
	// 		"qwer":  "2001",
	// 		"qwert": "2002",
	// 	},
	// 	"Mary": {},
	// 	"Paul": {
	// 		"zxc": "2003",
	// 		"cxz": "2004",
	// 	},
	// }
	// for a, k := range author {
	// 	if len(k) == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(len(author[a]))
	// 	break
	// }
	// fmt.Println(author)

}
