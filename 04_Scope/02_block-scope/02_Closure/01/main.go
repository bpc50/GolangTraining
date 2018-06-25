package main

import "fmt"

func main() {
	x := 42 //Anything declared in the outer scope, is available anywhere within the inner scope below.
	fmt.Println(x)
	{
		fmt.Println(x + 1) //The x var was declared in the outer scope, so is available within the inner scope.
		y := "Going to Maine soon, yay."
		fmt.Println(y) //The x var is declared in the inner scope, so it's not available in the outer scope.
	}
	//fmt.Println(y) //If you tried to call y outside of the inner scope, it would error.
}
