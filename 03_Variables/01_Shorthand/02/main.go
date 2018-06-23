package main

import "fmt"

func main() {
	var e string
	e = "Bikes are cool"

	//Declaring AND Assigning a value to it straight from the beginning.
	a := 11
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%T \n", a) //%T is going to show the type that go decides your variable is.
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
}
