package main

import (
	"errors"
	"os"
)

func main() {
	a
	if _, err := os.Open("myFile.txt"); err != nil {
		if errors.Is(err, os.ErrNotExist)
	}
}
