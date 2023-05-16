package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestClosure(t *testing.T) {
	name := "gandhi"
	counter := 0

	increment := func() {
		name := "setyawan"
		fmt.Println("increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
