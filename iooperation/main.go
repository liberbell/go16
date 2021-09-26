package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {

	//Make directory
	d, err := os.Stat("subdir")
	fmt.Println("error returned by os.Stat is: ", err)

	if err != nil {
		fmt.Println(d.Mode(), d.IsDir())
		log.Fatal("file/directory name already exists.")
	}
	if errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll("subdir", 0777)
	}
}
