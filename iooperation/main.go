package main

import (
	"os"
	"path/filepath"
)

func main() {

	//Make directory
	// d, err := os.Stat("subdir")
	// fmt.Println("error returned by os.Stat is: ", err)

	// if err == nil {
	// 	fmt.Println(d.Mode(), d.IsDir())
	// 	log.Fatal("file/directory name already exists.")
	// }
	// if errors.Is(err, os.ErrNotExist) {
	// 	err := os.MkdirAll("subdir", 0777)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("subdir directory created.")
	// }

	//Make directory all
	p := filepath.Join("../test", "subdir1", "subdir2")
	err := os.MkdirAll(p, 0777)
}
