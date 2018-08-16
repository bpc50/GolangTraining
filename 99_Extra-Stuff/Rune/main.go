package main

import "fmt"

func main() {
	for i := 47; i <= 70; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}
