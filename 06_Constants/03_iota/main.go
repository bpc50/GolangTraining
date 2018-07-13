package main

import "fmt"

const (				//iota starts at 0, and auto advances by 1 with each use.
	A		= iota
	B		= iota
	C		= iota
)

const (				//each time it's redefined in a new const block, it restarts at 0.
	D		= iota
)

const (				//can also declare/advance using iota like this.  works same as above in a, b, c.
	E		= iota
	F
	G
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
}