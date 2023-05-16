/*
* Struct
- template data yang di gunakan untuk menggabungkan nol atau lebih tipe data lainya dalam satu kesatuan
- struct biasanya representasi data dalam program aplikasi yg kita buat
- data di struct disimpan dalam field
- sederhanaya struct adalah kumpulan dari field

format :

	type NameStruct struct {
		NameField1, NameField2 typeDataField
	}

ex :

	type Customer struct {
		Name, Address string
		Age	int
	}

Membuat Dta Struct
-Struct : template date / prototipe data
-struct tidak dapat langsung di gunakan,namun dapat membuat data/object dari struct yang di buat
*/
package learn_golang_basic

import (
	"fmt"
	"testing"
)

type Customer struct {
	Name, Address string
	Age           int
}

func TestStruct(t *testing.T) {
	var result Customer
	result.Name = "gandhi"
	result.Address = "indonesia"
	result.Age = 28

	fmt.Println(result)
	fmt.Println(result.Name)
}
