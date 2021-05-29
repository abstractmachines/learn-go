package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Const, immutable.
	const country = "Ireland"
	fmt.Println(country)
	// country = "new value that will fail due to const"

	fmt.Println("wut up G.")
	fmt.Println(600 + 442)
	fmt.Println("I love" + "go")
	fmt.Println(false || true)

	// Var, mutable.
	var age int

	fmt.Println("age is: ", age)
	fmt.Println(reflect.TypeOf(age))

	var name = "squib"
	/* equivalencies: either have the TYPE or the assignment operator:
	var name string
	var name = "squib"
	*/

	fmt.Println("name is: ", name)
	fmt.Println(reflect.TypeOf(name))

	shortDeclare := "crazyFoo"

	fmt.Println(shortDeclare)
}
