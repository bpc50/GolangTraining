package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(y) //Why is y callable even though it appears to not have been declared yet?
	//Because it's outside of the curly-braces.  Anything outside processes 1st, and then is available to
	//  anything instead of curly braces/bracket, etc.
}

var y = 52
