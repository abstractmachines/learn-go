package main

import "fmt"
import "reflect"

func main() {
	var one = 1
	var two = 2
	var resultInt = one * two
	fmt.Println(resultInt)
	fmt.Println(reflect.TypeOf(resultInt)) // int.

	/*
		Question: Can I do an implicit type conversion in golang?
		Answer: Nope, you can't (unless it's an untyped constant, more later).
	*/
	// The compiler will catch this code. "mismatched types."
	//var three = 3.0 // float
	//var resultMultiPass = three * one
	//fmt.Println(resultMultiPass)

	/* So in golang you have to have an explicit type conversion.
	 */
	
}