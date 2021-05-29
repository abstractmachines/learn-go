package main

import (
	"fmt"
)

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

func main() {
	var buffer [256]byte

	type sliceHeader struct {
		Length        int
		ZerothElement *byte
	}

	// ** this just prints out hex memory for value of slice.
	//slice := sliceHeader{
	//	Length:        50,
	//	ZerothElement: &buffer[100],
	//}
	//fmt.Println("slice: ", slice)

	slice := buffer[10:20]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
	fmt.Println("before", slice)
	AddOneToEachElement(slice)
	fmt.Println("after", slice)
}
