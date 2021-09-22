package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// signature: func open
	f, err := os.Open("myFile.txt")
	defer f.Close()

	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("file successfully opened.", f.Name())
}
