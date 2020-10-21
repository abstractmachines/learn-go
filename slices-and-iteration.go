package main

import "fmt"

/*
Arrays are limiting because the SIZE is part of the TYPE.
Arrays really aren't common in golang.

We iterate over the array here with a simple for loop.
 */
func anArray() [5]int {
	var intArray [5]int

	var lengthOf = len(intArray)

	var val int = 0
	for i := 0; i < 5; i++ {
		intArray[i] = val
		fmt.Println(intArray[i])
		val++
	}

	fmt.Printf("\n\nLength of array of ints is %v \n", lengthOf)

	return intArray
}

/*
- For Slice and Map data structures, we iterate with a "range" variation of the for loop.

- When ranging over a slice, two values are returned for each iteration.
The first is the index, and the second is a copy of the element at that index.
 */
func sliceIt() {
	type sliceHeader struct {
		Length int
		ZerothIndex *int
	}


}

func main() {
	var intArray [5]int = anArray()
	fmt.Println("this is the int array from main. ", intArray)
}
