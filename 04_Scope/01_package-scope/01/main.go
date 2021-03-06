package main

import "fmt"

var x int = 42 //Package level scope, since it's not declared within a func, x is available in the whole package.
//Remember, the Package is the Folder, so x is now available in any file added to this folder/package.

func main() {
	fmt.Println(x)

	foo()

	y := "BPCizzle" //Function level scope, defined and thus y only available in func main.
	fmt.Println(y)  //and in Function level scope, order matters, you can't call y to print, before you've declared it.
}

func foo() {
	fmt.Println(x)
	//	fmt.Println(y) //Since y has only Function level scope, it's not available in func foo, this would error.
}
