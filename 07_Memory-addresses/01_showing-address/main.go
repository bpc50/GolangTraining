package main

import "fmt"

func main() {

	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)	//ampersand shows the memory address as a hexadecimal.
	fmt.Printf("%d \n", &a)		//now let's convert the hexa into the decimal representation of the memory location, and show on a new line. 

}