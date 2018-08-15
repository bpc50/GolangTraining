package main

import "fmt"

func main() {
	i := 0				//initialize.
	for {				//oops, no condition, this will run forever (bad)
		fmt.Println(i)
		i++				//post, advance by 1.	
	}
}