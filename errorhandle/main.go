package main

import (
	"log"
	"os"
)

func main() {

	// signature: func open
	f, err := os.Open("myFile.txt")
	defer f.Close()

	if err != nil {
		log.Println(err)
	}
}
