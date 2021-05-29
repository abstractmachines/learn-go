package main

import (
	"encoding/json"
	"fmt"
)

// example from encoding/json package - json.Unmarshal() function

func main() {
	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal                      // slice of animal
	err := json.Unmarshal(jsonBlob, &animals) // Unmarshal jsonBlob into a slice of animals. input, output.
	/*
		Why do we need addressOf operator? Because of https://golang.org/pkg/encoding/json/#Unmarshal - function signature

		func Unmarshal(data []byte, v interface{}) error
		"Unmarshal ... stores the result in the value **pointed to by v.** " -> So, per the 2nd arg, we're writing to a pointer.
	*/
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}
