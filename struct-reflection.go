// For more info, see go101.org/article/reflection.html#struct-field-tag
package main

import (
	"fmt"
	"reflect"
)

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
