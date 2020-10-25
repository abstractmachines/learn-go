package main

import "fmt"

type S struct {
	x int
	y bool
}

func main() {
	var structUsage = S{8,false}
	var zeroValStruct = S{}
	var defaultValues = S{x:8}
	var orderDoesNotMatter = S{y:true, x:1}
	fmt.Println(structUsage, "\n", zeroValStruct, "\n", defaultValues, "\n", orderDoesNotMatter)
}
