package main

import "fmt"

func main() {
	i := 0				//initialize...the i variable.
	for i < 10 {		//condition...go doesn't have while, so this basically says, while i < 10, continue looping.
		fmt.Println(i)
		i++				//post...bump i up by 1 each time thru the loop, to help keep track of condition.	
	}
}