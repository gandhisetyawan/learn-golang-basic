package learn_golang_basic

import (
	"fmt"
	"testing"
)

func sayHello(firstName string, lastName string) {
	fmt.Println("hello,", firstName, lastName)
}

func TestFunction(t *testing.T) {
	sayHello("gandhi", "setyawan") // pemanggilan function
}
