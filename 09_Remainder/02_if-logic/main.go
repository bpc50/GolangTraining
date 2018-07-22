package main

import "fmt"

func main() {
	x := 13 % 3 // % is for Remainder, which would be 1.  We would of course use / for Division, which would be 4.
	fmt.Println(x)

	// Maybe we'd use this to print out if the first number is odd or even. Bad example, but, just using it here to show format of If
	//   conditional coding.
	if x == 1 {
		fmt.Println("Odd")
	} else { // Note the else has to be on the same line as the previous }.
		fmt.Println("Even")
	}
}
