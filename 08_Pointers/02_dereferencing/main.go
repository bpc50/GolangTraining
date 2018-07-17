package main

import (
	"fmt"
)

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a //Referencing the memory address here.  var b is of type 'int pointer', and is being assigned &a which is the memory address of a.

	fmt.Println(b)  //as in referencing, prints the memory address of var a which is now assigned to var b.
	fmt.Println(*b) //De-Referencing: The * removes the reference to memory address, and instead shows the actual value

	//De-referencing: Simply add the * operator in front of the variable that is holding the memory address, to see the value.
}
