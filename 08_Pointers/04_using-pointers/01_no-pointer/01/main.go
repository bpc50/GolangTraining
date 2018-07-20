package main

import "fmt"

func zero(z int) { // 2) x, now with a value of 5, is passed in, AS A COPY (not more efficiently pointing to memory), from below func main.
	z = 0 // 3) now we are changing that value of 5, to 0.  Not sure why, just because.
}

func main() {
	x := 5
	zero(x)        // 1) We are passing this x, up into the func zero function above.
	fmt.Println(x) // 4) But x is still 5.
}
