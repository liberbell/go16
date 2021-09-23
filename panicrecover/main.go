package main

import "fmt"

func p(s string, i int) {
	x := [3]int{1, 2, 3}
	x[i] = 11
	if i > 2 {
		panic(s)
	}
}

func r() {
	if err := recover(); err != nil {
		fmt.Println(err)
		fmt.Println("Recovered from panic.")
	}
}

func main() {
	defer r()
	p("runtime error: enter panic state.", 3)
}
