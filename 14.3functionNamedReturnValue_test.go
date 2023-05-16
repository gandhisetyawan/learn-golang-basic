package learn_golang_basic

import (
	"fmt"
	"testing"
)

func getFullNames() (firstName string, lastName string, age int) {
	firstName = "gandhi"
	lastName = "setyawan"
	age = 28
	return firstName, lastName, age
	// atau hanya menuliskan retur saja bisa
}

func TestFuncNamedReturnValue(t *testing.T) {
	a, b, c := getFullNames()
	fmt.Println(a, b, c)
}
