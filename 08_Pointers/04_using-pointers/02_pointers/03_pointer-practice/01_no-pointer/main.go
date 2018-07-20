package main

import "fmt"

func bpc(b int) { // 2) bpc is now passed up here, and then stored in new var b, as type int.
	b = b + 100
	fmt.Println(b) // 3) will be 200.
}

func main() {
	a := 100
	bpc(a) // 1) bpc is passed value of variable a.
	fmt.Println(a)	// 4) will be 100.
}
