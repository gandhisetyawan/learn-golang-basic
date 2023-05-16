/*
RECURSIVE FUNCTION
Recursive function : function yang memanggila dirinya sendiri
contoh : faktorial
*/
package learn_golang_basic

import (
	"fmt"
	"testing"
)

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		// x :=
		return value * factorialRecursive(value-1)
	}
}

func TestRecursiveFunction(t *testing.T) {
	fmt.Println(5 * 4 * 3 * 2 * 1)
	fmt.Println(factorialRecursive(5))
}
