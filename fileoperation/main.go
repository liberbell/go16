package main

import (
	"fmt"
	"os"
)

//Open and Close file
func closer(f *os.File) error {
	f.Close()
	fmt.Println(f.Name(), " successfully closed.")
}

func main() {
	//Create File
	// f, err := os.Create("create.txt")
	// defer f.Close()

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(f)

	//Open and Cloes file

}
