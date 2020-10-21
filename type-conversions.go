package main

import "fmt"
import "reflect"

func main() {
	var oneInt = 1
	var twoInt = 2
	var resultInt = oneInt * twoInt
	fmt.Println(resultInt)
	fmt.Println("typeOf result Int should be int: ", reflect.TypeOf(resultInt)) // int.

	/*
		Question: Can I do an implicit type conversion in golang?
		Answer: Nope, you can't (unless it's an untyped constant, more later).
	*/
	// The compiler will catch this code. "mismatched types."
	//var threeFloat = 3.0 // float
	//var resultMultiPass = three * one
	//fmt.Println(resultMultiPass)

	/* So in golang you have to CAST for an explicit type conversion.
	 */
	fmt.Println(reflect.TypeOf(oneInt))
	var threeFloat = 3.0 // could also be: var threeFloat float64 = 3.0 ... or threeFloat := 3.0
	resultFloat := float64(oneInt) * threeFloat
	fmt.Println(resultFloat)
	
}