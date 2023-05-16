/*
FUNCTION AS PARAMETER
Function tidak hanya bisa di simpan pada sebuah variabel.
Namun  juga bisa di gunakan sebagai parameter
*/
package learn_golang_basic

import (
	"fmt"
	"testing"
)

func greetingWithFilter(name string, filter func(string) string) {
	fmt.Println("hello", filter(name))
}

func spamFilter(name string) string {
	if name == "asu" {
		return "bicara yg baik"
	} else {
		return name
	}
}

func TestFuncAsParameter(t *testing.T) {
	greetingWithFilter("gandhi", spamFilter)

	filter := spamFilter // tampung di variabel filter untuk di jadikan argumen dari sebuah function
	greetingWithFilter("asu", filter)
}
