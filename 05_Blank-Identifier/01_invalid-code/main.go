package main

import "fmt"

func main() {
	a := "stored in a"
	b := "stored in b"
	//fmt.Println("a - ", a)		//won't run, because we've declared b as var, but didn't use it. 
	fmt.Println("a - ", a, ".  b - ", b)
}