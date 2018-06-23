package main

import "fmt"

func main() {
	//Best.  Shorthand variable declaration. Simple, done within a func, go will infer the type, no need to specify.
	//Declare, Initialize, & Assign, all at the same time.
	a := 11
	b := "golang"
	c := 4.17
	d := true

	//for example, this is how we'd explicitly declare, if need be:
	var e string
	e = "Motorcycles are cool"

	//or, similiarly but even more simply:
	var f string = "Bikes are awesome!"

	fmt.Printf("%v \n", a) //%v is "value in default format", \n is "next line", a is the variable of course.
	fmt.Printf("%v \n", b) //%v really figures alot of stuff out for you, ie, it figures out the type, and then outputs it properly
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)

}
