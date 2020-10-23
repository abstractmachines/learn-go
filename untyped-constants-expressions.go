package main

import "fmt"

func main() {
	/*
	We know that golang uses type inference for untyped constants; so it's a bit of a misnomer to call these variables
	"untyped", as each "untyped constant" in golang has an inherent, "default type." But for now we'll think of it as
	"variables that haven't had a type such as `string` or `int` or `int32` explicitly defined, are untyped."

	Since constants are always untyped in golang until the specific type is explicitly assigned, you can use those
	constants in expressions.
	 */

	// Constant Expressions.
	const a = 1 + 2.33
	fmt.Println(a) // 3.33
	const b = 'f' + 2
	fmt.Println(b) // ascii of f + 2 (104)

	const c = 1
	const d = 2.33
	const e = c + d
	fmt.Printf("the value is still the same, it's %v which is of type %T\n", e, e) // 3.33
}
