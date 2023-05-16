/** Struct Function atau Struct Method
- Struct merupakan tipe data seperti tipe data lainya, dia bisa di jadikan parameter untuk function
- struct tdk meiliki behavior ,
	namun dapat menambhaka method/function pada struct, sehingga seolah olah struct memiliki function
*/

package learn_golang_basic

import (
	"fmt"
	"testing"
)

type DaftarCustomer struct {
	Name, Address string
	Age           int
}

func (daftarCustomer DaftarCustomer) greeting(name string) {
	fmt.Println("Hello", name, "My Name is", daftarCustomer.Name, "alamat:", daftarCustomer.Address)
}

func TestStructFunction(t *testing.T) {
	var data DaftarCustomer = DaftarCustomer{
		Name:    "gandhi s",
		Address: "indonesia",
		Age:     28,
	}

	data.greeting("iwan")
}
