/*FUNCTION TYPE DECLARATIONS
-Terkadang jika function parameter terlalu panjang, maka agak ribet menulikan
di parameter sebuah function secara langsung.
-Type Declaration bisa di gunakan membuat alis function,
sehingga mempermudah menggunkan function sebagai parameter.
*/

package learn_golang_basic

import (
	"fmt"
	"testing"
)

type filter func(string) string

func greetingWithFilters(name string, filter filter) {
	fmt.Println("hello", filter(name))
}

// function yang di jadikan argument
func spamFilters(name string) string {
	if name == "asu" {
		return "bicara yg baik"
	} else {
		return name
	}
}

func TestFunctionTypeDeclaration(t *testing.T) {
	greetingWithFilters("gandhi", spamFilter)

	filter := spamFilters // tampung di variabel filter untuk di jadikan argumen dari sebuah function
	greetingWithFilter("asu", filter)
}
