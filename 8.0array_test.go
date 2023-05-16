package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var names [3]string
	names[0] = "gandhi"
	names[1] = "setyawan"
	names[2] = "iwan"

	//memberikan nilai langsung pada variabel
	var values = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//function array
	values[0] = 100
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(len(values))
}
