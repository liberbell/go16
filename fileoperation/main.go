package main

import (
	"fmt"
	"log"
	"os"
)

//Open and Close file
// func closer(f *os.File) error {
// 	f.Close()
// 	fmt.Println(f.Name(), " successfully closed.")
// 	return nil
// }

func main() {
	//Create File
	// f, err := os.Create("create.txt")
	// defer f.Close()

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(f)

	//Open and Cloes file
	// f, err := os.Open("create.txt")
	// defer closer(f)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("file successfully opened.", f.Name())

	//Remove delete file
	err := os.Remove("del.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file successfully removed.")
}
