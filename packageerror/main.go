package main

import (
	"errors"
	"log"
	"os"
)

func main() {

	if _, err := os.Open("myFile.txt"); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Println("File does not exist.")
		} else {
			log.Println(err)
		}
		return
	}

}
