package main

import "fmt"

func bpc(b *int) { // 2) bpc is now passed up here, and then stored in new var b, as type int.
	*b = 200
	fmt.Println(*b) // 3) will be 200.
}

func main() {
	a := 100
	bpc(&a)        // 1) bpc is passed memory address (which remember, also holds the value, though not shown) of variable a.
	fmt.Println(a) // 4) will ALSO be 200, cause changing the value, after dereferencing from memory address, changes orig var value
}
