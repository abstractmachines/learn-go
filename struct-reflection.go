// For more info, see go101.org/article/reflection.html#struct-field-tag
package main

import (
	"fmt"
	"reflect"
)

/* Structs and Field Tags

Like C, Go has structs (with extra bonuses such as optional Field Tags following the field name and data type).

The Field Tag of a struct is usually a key-value pair that communicates meta information about the field.

The default value for a struct's field tag is an empty string (if not explicitly declared).
 */
type Structy struct {
	AField string `max: "99"`
	AnotherField int `optional: "yes"`
}

func main() {
	t := reflect.TypeOf(Structy{})
	fmt.Println("\nt: ", t)

	x, y := t.Field(0).Tag, t.Field(1).Tag
	fmt.Println(reflect.TypeOf(x)) // reflect.StructTag
	fmt.Println(reflect.TypeOf(y)) // reflect.StructTag


}
