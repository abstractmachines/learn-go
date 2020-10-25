package main

import "fmt"
//import "reflect"

func main() {
	// integer variables are mutable
	oneNumber := 1
	oneNumber = 2
	fmt.Println(oneNumber)

	// string variables are mutable
	oneString := "hey"
	oneString = "sup"
	fmt.Println(oneString)

	/*
	https://stackoverflow.com/questions/8018081/which-types-are-mutable-and-immutable-in-the-google-go-language

	Mutable Go objects:

	arrays and slices
	maps
	channels
	closures which are capturing at least 1 variable from the outer scope
	Immutable Go objects:

	interfaces
	booleans, numeric values (including values of type int)
	strings
	pointers
	function pointers, and closures which can be reduced to function pointers
	*/
}
