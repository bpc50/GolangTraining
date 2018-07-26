package main

import "fmt"

func main() {
	x := 13 % 3 // % is for Remainder, which would be 1.  We would of course use / for Division, which would be 4.
	fmt.Println(x)

	// Maybe we'd use this to print out if the first number is odd or even. Bad example, but, just using it here to show format of If
	//   conditional coding.  Ex, i is 4. The remainder (%) of 4 / 2 is 0, ie, Else.
	for i := 1; i < 10; i++ {
		if i%2 == 1 { // Note, == means 'is equal to'..., whereas = is an assignment.
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}
}
