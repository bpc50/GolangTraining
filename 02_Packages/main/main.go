package main

import (
	"fmt"

	"github.com/bpc50/GolangTraining/02_Packages/stringUtil"
)

func main() {
	fmt.Println(stringUtil.MyName) //grab the folder/package "stringUtil", and the variable "MyName" from one of the files (name.go) within it.
}
