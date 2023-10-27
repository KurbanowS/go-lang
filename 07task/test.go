package main

import (
	"log"
	"os"
)

func main() {
	// -----------------------------------------------------------
	myfile, e := os.Stat("data/in.txt")
	if e != nil {

		// Checking if the given file exists or not
		// Using IsNotExist() function
		if os.IsNotExist(e) {
			log.Fatal("File not Found !!")
		}
	}
	log.Println("File Exist!!")
	log.Println("Detail of file is:")
	log.Println("Name: ", myfile.Name())
	log.Println("Size: ", myfile.Size())

	// ---------------------------------------------------------

	// sourceFile := "src.txt"
	// destinationFile := "dst.txt"
	// source, err := os.Open(sourceFile) //open the source file
	// if err != nil {
	// 	panic(err)
	// }
	// defer source.Close()
	// destination, err := os.Create(destinationFile) //create the destination file
	// if err != nil {
	// 	panic(err)
	// }
	// defer destination.Close()
	// _, err = io.Copy(destination, source) //copy the contents of source to destination file
	// if err != nil {
	// 	panic(err)
	// }

}
