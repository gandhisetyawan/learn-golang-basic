/*
  - Pointer Di function
  - secara default membuat parameter di function adalah pass by value
    ( data di copy lalu di kirim ke function)
  - ketika mengubah data di function maka data asli tidak ikut berubah
  - Namun kadang kita memerlukan function yang dapat meruba data aslinya
  - untuk melakukan function dapat merubah data aslinya , maka gunakan pointer
  - untuk menjadikan parametr menjadi pointer tambahkan tanda * pada parameternya
*/
package learn_golang_basic

import (
	"fmt"
	"testing"
)

type Alamat struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Alamat) {
	address.Country = "Indonesia"
	fmt.Println(address)
}

func TestPointerFunction(t *testing.T) {
	var address Alamat = Alamat{
		City:     "Cikarang",
		Province: "Jawa Barat",
		Country:  "",
	}
	var addressPointer *Alamat = &address
	ChangeAddressToIndonesia(addressPointer)

	fmt.Println(address)
}
