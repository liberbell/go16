package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("create.txt")
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
}
