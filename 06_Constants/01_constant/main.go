package main

import "fmt"

//const p string = "death & taxes"	//you can write this either way, ie, implicitly declaring the type, ie, string...or not. 
const p = "death and taxes"

//const q = 42

func main() {

		const q = 42	//Notice when declaring constant variable, it's =.  Reg variable assignment is :=.

		r := "death2 and taxes2"

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println("r - ", r)
}

//a CONSTANT is a simple unchanging value
