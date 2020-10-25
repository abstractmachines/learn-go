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
	var animals []Animal // slice of animal
	err := json.Unmarshal(jsonBlob, &animals) // Unmarshal jsonBlob into a slice of animals. input, output.
	// address of operator to marshal into a slice ..
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}