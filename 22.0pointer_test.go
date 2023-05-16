/*
* POINTER
-secara default di golang semua variabel di passing by value , bukan by reference
-artinya , jika mengirimkan sebuah variabel ke dalam function,method atau variabel lain,
sebenarnya yang dikirim adalah duplikasi/copy valunya

Pointer :
-pointer kemampuan membuat reference ke lokasi data di memori yang sama, tanpa menduplikasi
data yang sudah ada
-gampangnya , dengan pointer kita dapat membuat pass by reference
*/
package learn_golang_basic

import (
	"fmt"
	"testing"
)

type Address struct {
	City, Province, Country string
}

func TestPointer(t *testing.T) {
	var data1 Address = Address{
		City:     "cikarang",
		Province: "jawabarat",
		Country:  "indonesia",
	}
	/*tidak reference ke data1, namun data2 copy/duplikasi dari data1
	jadi ketika data2 ada perubahan data1 tidak ikut berubah
	*/
	var data2 Address = data1

	/*Pointer : data3 reference ke data1
	jadi ketika ada perubahan di data3 makan data1 juga akan berubah mengikuti
	*/
	// data3 := &data1
	var data3 *Address = &data1 // pointer, data3 reference ke data1

	//edit field
	data2.City = "Tulungagung" //perubahan data2 , tidak ngaruh ke field City data1
	data3.City = "Bekasi"      // perubahan data3 pengaruh ke field City data1

	/*Tanpa Operator *
	inisialisasi ulang variabel pointer data3 tanpa operator *
	maka yang berubah hanya variabel pointer data3
	*/
	data3 = &Address{
		City:     "Malang",
		Province: "jawaTimur",
		Country:  "indonesia",
	}

	fmt.Println(data1)
	fmt.Println(data2)
	fmt.Println(data3)

}
