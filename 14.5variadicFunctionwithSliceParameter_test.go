package learn_golang_basic

import (
	"fmt"
	"testing"
)

func sumAllNumber(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func TestVariadicFuncWithSliceParameter(t *testing.T) {
	total := sumAllNumber(10, 10, 10, 10)
	fmt.Println(total)

	//parameter nya berupa tipe data slice
	numbers := []int{10, 10, 10}
	total = sumAllNumber(numbers...)
	fmt.Println(total)
}
