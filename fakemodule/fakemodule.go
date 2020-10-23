package fakemodule

import "fmt"

func Fakemodule() string {
	//message: := fmt.Sprintf("Wut up %v, sup dude? ", name)
	var message = fmt.Sprint("hey dude")
	return message
}
