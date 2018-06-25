package main

import "fmt"

func wrapper() func() int { //wrapper parameters would go in teh (), and here func() and int is the "return". So we are returning a function.
	x := 0
	return func() int { //and here is the function we are returning
		x++
		return x
	}
}

//*So when wrapper gets called, the func() defined after "return" is what gets returned.
// And then "wrapper" gets assigned to increment below.

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
