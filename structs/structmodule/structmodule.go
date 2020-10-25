// For more info, see go101.org/article/reflection.html#struct-field-tag
package structmodule

import "fmt"

/*
We have "encapsulation" via non-exported struct fields. We consider those "private."

How to export identifiers in Go: Capitalize it. That will inherently export that identifier.
 */

func Structmodule() string {
	//type StructyPublicCapitalized struct {
	//	PublicField string
	//	privateField string
	//}
	//
	//fmt.Println(StructyPublicCapitalized{})
	//
	//babyT := StructyPublicCapitalized{
	//	PublicField: "I am a Public struct field",
	//	privateField: "I am a Private struct field",
	//}
	//fmt.Println(babyT)
	var message = fmt.Sprint("hey dude")
	return message
}
