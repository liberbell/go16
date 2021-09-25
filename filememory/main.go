package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {

	contents, err := os.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reflect.TypeOf(contents))
	fmt.Println(contents)
	fmt.Println(string(contents))

}
