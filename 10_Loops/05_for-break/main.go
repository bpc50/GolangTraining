package main

import "fmt"

func main() {
	i := 0					//initialize.
	for {
		fmt.Println(i)
		if i >= 20 {		//condition. this would be the same as the 04 infinite loop (bad), but the if/condition/break saves it. 
			break
		}
		i++					//post.  
	}
}