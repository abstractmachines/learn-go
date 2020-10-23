// For more info, see go101.org/article/reflection.html#struct-field-tag
package main

import "fmt"

/*
We have "encapsulation" via non-exported struct fields. We consider those "private."

How to export identifiers in Go: Capitalize it. That will inherently export that identifier.
 */

type StructyPublicCapitalized struct {
	CanHazExportCapitalized string
	izPrivateLowercase string
}

func main() {
	fmt.Println(StructyPublicCapitalized{})

	babyT := StructyPublicCapitalized{
		CanHazExportCapitalized: "I am public",
		izPrivateLowercase: "I am private",
	}
	fmt.Println(babyT)
}
