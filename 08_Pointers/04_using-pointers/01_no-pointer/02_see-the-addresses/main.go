package main

import "fmt"

func zero(z int) { // 1) We are passing over the value 5 from zero below, into this zero, and then storing it in z.
	fmt.Println(&z)
	z = 0 // 2) Now we are changing z, from 5, to 0.
}

func main() {
	x := 5
	fmt.Println(&x)
	zero(x)
	fmt.Println(x) // 3) but x is still it's own thing.
}
