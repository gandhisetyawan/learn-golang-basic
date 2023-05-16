package learn_golang_basic

import (
	"fmt"
	"testing"
)

func getFullName(firstName string, lastName string, age int) (string, string, int) {
	return firstName, lastName, age
}

func TestFuncReturnMultipleValues(t *testing.T) {
	firstName, lastName, age := getFullName("gandhi", "setyawan", 28)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
	fmt.Println(getFullName("yurita", "tri", 26))
}
