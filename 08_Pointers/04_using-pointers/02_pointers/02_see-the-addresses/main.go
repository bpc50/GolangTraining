package main

import "fmt"

func zero(z *int) { // 2) get value of zero from below, ie, the memory address of 5, store that as new variable z, type pointer int.
	fmt.Println(z)
	*z = 0 // 3) deference memory address, meaning, turn it back to a value.  Of 0...which will auto pass back to x below.
}

func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)       // 1) store the memory address of 5, in zero, to pass to above func zero.
	fmt.Println(x) // 4) final value is now 0.
}
