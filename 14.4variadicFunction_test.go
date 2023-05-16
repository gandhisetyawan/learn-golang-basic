package learn_golang_basic

import (
	"fmt"
	"testing"
)

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total
}

func TestVariadicFunction(t *testing.T) {
	total := sumAll(10, 10, 10)
	fmt.Println(total)
}
