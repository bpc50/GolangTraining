package main

import "fmt"

func main() {

	//Declaring/Assigning variable an initial value of 0.  Ie, when you don't specify, Go will automatically assign a starting value.
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%v \n", a) //int initial value will be 0
	fmt.Printf("%v \n", b) //string initial value will be a blank
	fmt.Printf("%v \n", c) //float64 initial value will be a 0
	fmt.Printf("%v \n", d) //boolean initial value will be false
}
