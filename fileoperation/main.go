package main

//Open and Close file
// func closer(f *os.File) error {
// 	f.Close()
// 	fmt.Println(f.Name(), " successfully closed.")
// 	return nil
// }

func main() {
	//Create File
	// f, err := os.Create("create.txt")
	// defer f.Close()

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(f)

	//Open and Cloes file
	// f, err := os.Open("create.txt")
	// defer closer(f)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("file successfully opened.", f.Name())

	//Remove delete file
	// err := os.Remove("del.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("file successfully removed.")

	//Copy file
	// src, err := os.Open("src.txt")
	// defer src.Close()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// dst, err := os.OpenFile("dst.txt", os.O_RDWR|os.O_CREATE, 0755)
	// defer dst.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// w, err := io.Copy(dst, src)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(reflect.TypeOf(w))
	// fmt.Println(w)

	//Rename file

}
