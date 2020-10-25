package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// example from encoding/json package - json.Marshal() function

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"}, // Composite Literal
	}
	b, err := json.Marshal(group) // dot notation. We are catching "b" and "err." In Go, we check errors right after where they'd be returned.
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}