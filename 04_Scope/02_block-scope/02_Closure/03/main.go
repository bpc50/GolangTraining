package main

import "fmt"

func main() {
	x := 0                     //declaring the variable within a func is good, cause it's narrowing the scope down.
	increment := func() int {  //Here we are assigning a function, to a variable (called "func expression". The func is called an anonymous func.
		x++ 				   //The ++ is a way to say "add 1 to the variable"
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
