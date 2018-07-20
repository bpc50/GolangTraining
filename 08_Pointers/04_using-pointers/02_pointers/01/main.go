package main

import "fmt"

func zero(z *int) { // 3) we pass over the memory address, to zero, and now store it in a new variable z, as type "pointer to an int (*int)"
	*z = 0 // 4) now I "dereference it z using the * command, changing it the actual value of z to a 0"
}

func main() {
	x := 5         // 1) x is 5
	zero(&x)       // 2) we store the memory address of 5, in zero
	fmt.Println(x) // 5) 4 above, automatically changes x, to 0.  (Confusing at this point, for sure.)
}
