package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
}
