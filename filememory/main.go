package main

import (
	"log"
	"os"
)

func main() {

	contents, err := os.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}

}
