package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")

	//Scan will sit and wait for User Answer, once recieved, it'll deliver answer, as meters, and continue formula. 
	fmt.Scan(&meters)	//Accept User Input.  Just like Print writes to standard out, Scan recieves from standard in.
						//Said more simply: When someone writes something with their keyboard, catch it!
						//Not in Go Doc. Just have to know.

	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}