package main

import "fmt"

func main() {
	for i := 47; i <= 70; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i))) //Rune is a character, and a string is a collection of Rune's.
	}

	foo := 'a' //foo is initialized as the string (ASCII) value of 'a', which is 97.  Different than "a", which is literal value, ie, type string.
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)       //print the type.
	fmt.Printf("%T \n", rune(foo)) //print the type, after converting the foo type to a rune. rune is just an alias for int32, so it's the same.
}

//Rune:  above we used "a slice of byte, ie, []byte(string(i))" to print the Rune value of i.
//So in each loop iteratio we printed 1) loop counter value 2) the string value of that 3) the rune value of that.
