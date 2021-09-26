package main

import (
	"fmt"
	"os"
)

func main() {

	//Make directory
	d, err := os.Stat("subdir")
	fmt.Println("error returned by os.Stat is: ", err)

	if err != nil {
		fmt.Println(d.Mode(), d.IsDir())
	}
}
