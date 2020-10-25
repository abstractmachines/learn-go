package main

import (
	"fmt"
	"structmodule"
)

func main() {
	message := structmodule.Structmodule()
	fmt.Println(message)

	fmt.Println(structmodule.Structy{})
	//
	babyT := structmodule.Structy{
		PublicField: "I am a Public struct field",
		//privateField: "I am a Private struct field", -> This won't compile, as it's a private field, as expected!
	}
	fmt.Println(babyT)
}


