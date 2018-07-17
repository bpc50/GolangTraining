package main

import (
	"fmt"
)

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 42 //the value at this address, change it to this.

	fmt.Println(a) //and now, since we changed the value at the memory address of a	, the value of a is also changed.

	//Why use this?
	// We can pass a memory address, instead of a bunch of values. (we can pass a reference)
	// Then, we can still change the value of whatever is stored in that memory address.
	// This makes our programs more performant.
	//  We don't have to pass around large amounts of data.
	//	We only have to to pass around the addresses.

	//In Go, everything is PASS BY VALUE, ie, when we pass a memory address, we pass the value as well.
}
