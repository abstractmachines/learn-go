package main

import (
	"fmt"
)

func main() {
	// string variables are immutable ...
	oneString := "hey"
	oneString = "sup" // This is new data.
	fmt.Println(oneString)
}
