package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	path := "data/in.txt"
	path2 := "data/out.txt"

	getFile(path, path2)

}
func getFile(path string, path2 string) {
	open, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	// read := bufio.NewScanner(open)
	// for read.Scan() {
	// 	fmt.Println(read.Text())
	// }
	defer open.Close()
	_, e := os.Stat(path2)
	if e != nil {
		if os.IsNotExist(e) {
			// log.Fatal("File not Found !!")
			destination, err := os.Create(path2)
			if err != nil {
				panic(err)
			}
			defer destination.Close()
			_, err = io.Copy(destination, open)
			if err != nil {
				panic(err)
			}
			fmt.Println(4 + 6)
		} else {
			destination, err := os.Open(path2)
			if err != nil {
				panic(err)
			}
			defer destination.Close()
			_, err = io.Copy(destination, open)
			if err != nil {
				panic(err)
			}
		}
	}
}
