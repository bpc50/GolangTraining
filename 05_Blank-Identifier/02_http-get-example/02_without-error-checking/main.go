package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	//Blank identifier _ put in place of err in this case. Put as a place holder when not ready to use a variable, but might later.
	//  Most of the time used like this when we are returning a value to a variable. 
	res, _ := http.Get("http://www.allsearchinc.com/")	//do http.Get, to get a result, put that into variable we'll call res.
	page , _ := ioutil.ReadAll(res.Body)	//do a ReadAll to read the responses Body, store in a variable we'll call page.
	res.Body.Close()	//then we have to always Close the response body.
	fmt.Printf("%s", page)	//then we'll print the page variable which now holds all the code from the web page, as a string ("%s")
}