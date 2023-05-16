/*
* Type Assertions
- Type assertion : kemampuan merubah tipe data menjadi tipe data yang di inginkan
- Fitur ini sering di gunakan ketika bertemu dengan data interface kosong
*/
package learn_golang_basic

import (
	"fmt"
	"testing"
)

func random() interface{} {
	return "ada saya string"
}

func TestTypeAssertions(t *testing.T) {
	result := random()
	resultString := result.(string) //konversi dari tipe data interface ke string
	fmt.Println(resultString)

	// cek tipe data kembalian : result.(type)
	switch value := result.(type) {
	case string:
		fmt.Println("value", value, "is string")

	case int:
		fmt.Println("value", value, "is integer")
	default:
		fmt.Println("unknown")
	}
	// ex jika terjadi kesalahan, panic
	resultInt := result.(int) // panic karena return data di interface string
	fmt.Println(resultInt)
}
