package main

import "fmt"

func main() {
	for i := 0; i < 2; i++ { //This is outside.
		for j := 0; j < 4; j++ { //This is nested, inside. Each time outside runs once, this runs thru all conditions.
			fmt.Println("Outer", i, " - ", "Inner", j)
		}
	}
}
