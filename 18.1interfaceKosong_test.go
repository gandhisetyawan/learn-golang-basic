/*
*
Interface Kosong
- Go-lang bukan bahasa pemrograman yang berorientasi object
- biasanya dalam pemrogram berorientasi object, ada satu data perent di puncak yang bisa
dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
- untuk menangani ini go-lang pakai interface kosong
- interface kosong : interface yang tidak memiliki deklarasi method satupun,
hal ini membuat secara otomatis semua tipe data akan menjadi implementasi nya
*/
package learn_golang_basic

import (
	"fmt"
	"testing"
)

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func InputTest(input interface{}) interface{} {
	if input == 1 {
		return 1
	} else if input == true {
		return true
	} else if input == false {
		return false
	} else {
		return "yang di input string"
	}
}

func TestInterfaceKosong(t *testing.T) {
	data1 := Ups(2) //var data1 interface{} = Ups()
	data2 := InputTest(1)

	fmt.Println(data1)
	fmt.Println(data2)
}
