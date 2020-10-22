package main

import "fmt"

/*
Arrays are limiting because the SIZE is part of the TYPE.
Arrays really aren't common in golang.

YOU HAVE TO size these arrays in these args or there's a panic runtime error "slice bounds out of range [:3] with capacity 0
^ I bet I'm missing something about dynamic arrays, but ... remember, Go uses mostly slices, not arrays.
 */
func createIntArray() [5]int {
	var intArray [5]int

	var val int = 0
	for i := 0; i < len(intArray); i++ {
		intArray[i] = val
		val++
	}

	fmt.Printf("\ncreateIntArray(): the newly created array is %v", intArray)
	return intArray
}

/*
- For Slice and Map data structures, we iterate with a "range" variation of the for loop.

- When ranging over a slice, two values are returned for each iteration.
The first is the index, and the second is a copy of the element at that index.
 */
func sliceStruct(arrayInt [5]int) {
	type sliceHeader struct {
		Length int
		Index *int
	}

	aSlice := sliceHeader {
		Length: len(arrayInt),
		Index: &arrayInt[1], // shows just a memory location. ok, broken. TODO fix it.
	}
	fmt.Println("Hey, so we sliced your array within a function at first index: ", aSlice)
}

func basicSlice(arr [5]int) []int {
	var slice1 []int = arr[1:3] // setting a capacity cap() of elements 1->5, so cap() === 4.
	return slice1
}

//func singleElementSlice(arr [5]int) int {
//	var
//}

func main() {
	var intArray [5]int = createIntArray()

	/*
	Let's perform a basic slice, and get accustomed to the slice API / methods. len() cap() etc
	 */
	var slice1 = basicSlice(intArray)

	// len() : length of a slice
	fmt.Printf("\nWe sliced arr[1:3], so new array is %v of len %v", slice1, len(slice1))

	// cap() : capacity of a slice
	fmt.Println("\ncap(slice) is ", cap(slice1))

	/*
	Slice to get one element of an array.
	 */
	//sliceSingle := intArray[:2]
	fmt.Println("array[:2] is : ", intArray[:2], " No index, and prints a non-inclusive range (cap).")
	fmt.Println("array[2:] is : ", intArray[2:], " Index/offset is at 2, and no range, so print cap from index offset.")

	// "Reslice our slice and extend it:"
	// let's use the capacity cap() of a slice to extend our slice to the full length of the array (AFTER our index into array).
	var sliceReslice = slice1[:cap(slice1)]
	fmt.Println("Extend! Reslice by using [:cap(slice)] and new array is ", sliceReslice) // [0, 1, 2, 3, 4]

	/*
	Slice an int literal.
	 */
	var sliceALiteral [] int = []int{5,6,7,8,9}
	fmt.Println("sliceALiteral ", sliceALiteral[:1])

	/*
	"Make" a slice.
	Make zeros out the values.
	 */
	makeSlice := make([]int, 5)
	fmt.Println("Make() a new slice: ", makeSlice)
	fmt.Printf("Type of slice looks like: %T\n", makeSlice)
	fmt.Printf("Type of array looks like: %T\n", intArray)

	/*
	Append to a slice.
	 */
	appendedSlice := append(slice1, 101)
	fmt.Printf("Slice: %v And append to it: %v\n", slice1, appendedSlice)
}
