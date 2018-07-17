package main

import (
	"fmt"
)

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a //variable b is of type "int pointer", ie pointing to a memory address where an int is stored.

	fmt.Println(b, "- is the value of variable b, ie, the memory location, in hexadecimal, of variable a")

	//above makes b a pointer to the memory address where an int is stored
	//assigning the memory address of var a, to var b.
	//Later we'll need to use var b, in a function.
}
