package main

import "fmt"

/*
Arrays are limiting because the SIZE is part of the TYPE.
Arrays really aren't common in golang.

We iterate over the array here with a simple for loop.
 */
func anArray() [5]int {
	var intArray [5]int

	var val int = 0
	for i := 0; i < len(intArray); i++ {
		intArray[i] = val
		fmt.Println(intArray[i])
		val++
	}

	return intArray
}

/*
- For Slice and Map data structures, we iterate with a "range" variation of the for loop.

- When ranging over a slice, two values are returned for each iteration.
The first is the index, and the second is a copy of the element at that index.
 */
func sliceIt(arrayInt [5]int) {
	type sliceHeader struct {
		Length int
		Index *int
	}

	aSlice := sliceHeader {
		Length: len(arrayInt),
		Index: &arrayInt[1], // shows just a memory location. ok, broken. TODO fix it.
	}
	//fmt.Println("this is arrayInt in sliceIt before structy tymes: ", arrayInt)
	fmt.Println("Hey, so we sliced your array within a function at first index: ", aSlice)
}

func main() {
	var intArray [5]int = anArray()
	fmt.Println("this is the int array from main. ", intArray)

	sliceIt(intArray)
}
