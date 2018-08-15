package main

import "fmt"

func main () {
	i := 0				//initialize.
	
	for {
		i++				//post, incrementer, counter, people call it different things.
		if i%2 == 0 { 	//condition.  Here we are saying if the remainder of i/2 is equal to 0, ie, an even #, continue to post/incrementor.
			continue
		}
		fmt.Println(i)  //we'd only land down here, if i%2 >< 0, ie was odd. In that case the if above wouldn't have executed continue.
		if i >= 50 {
			fmt.Println("Done")
			break
		}
	}
}