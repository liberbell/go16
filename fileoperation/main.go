package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//Create File
	f, err := os.Create("create.txt")
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
}
