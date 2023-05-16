/*TIPE DATA SLICE
- tipe data slice adalah potongan dari data array
- mirip dengan array, yg beda slice bisa berupah ukuranya
- slice dan array selalu terkoneksi,
	dimana slice adalah data yang mengakses sebagian atau seluruh data array
*/

/*
DETAIL TIPE DATA SLICE

+Detail Tipe Data Slice

	tipe data slice memiliki 3 data  : pointer , length , capacity
	- pinter	: penunjuk data pertama pada array para slice ( di slice)
	- length	: panjang data dari slice (yang di tampilkan )
	- capacity	: kapasitas dari slice , length tdk boleh lebih dari capacity

+Membuat slice dari Array
Perintah				Keterangan
----------------------------------------------------------------------------------------------------
array[low:high]			Membuat slice dari array di mulai dari index ke low sampai index sebelum high
array[low:]				Membuat slide dari array di mulai dari index low sampai index akhir di array
array[:high]			Membuat slice dari array di mulai dari index 0 sampai index sebelum high
array[:]				Membuat slice dari array di mulai dari index 0 sampai index akhir di array

+Function Slice
Operasi							Keterangan
----------------------------------------------------------------------------------------------
len(slice)						Untuk mendapatkan panjang(length)
cap(slice)						Untuk Mendapatkan kapasitas(capacity)
append(slice)					Membuat slice baru dengan menambah data ke posisi terakhir slice,Jika kapasitas penuh , maka akan membuat array baru

make(TypeData,length,capasity)	Membuat slice baru
copy(destination, source)		menyalin slice dari source ke destination
*/

package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestTipeDataSlice(t *testing.T) {
	//ARRAY
	var months = [...]string{
		"Januari",   //Index 0
		"Februari",  //Index 1
		"Maret",     //Index 2
		"April",     //Index 3
		"Mei",       //Index 4
		"Juni",      //Index 5
		"Juli",      //Index 6
		"Agustus",   //Index 7
		"September", //Index 8
		"Oktober",   //Index 9
		"November",  //Index 10
		"Desember",  //Index 11
	}

	//MEMBUAT SLICE DARI ARRAY
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println("Panjang slice : ", len(slice1))
	fmt.Println("Kapasitas slice : ", cap(slice1))

	//MENGUBAH DATA SLICE DAN DATA ARRAY
	slice1[0] = "may"
	fmt.Println(slice1)
	fmt.Println(months)

	/*FUNCTION APPEND*/
	fmt.Println("================>>> function append <<<============")
	var slice2 = months[10:]

	var slice3 = append(slice2, "data_baru") //tambah data slice, dan slice membuat array baru
	slice3[1] = "desember_baru"              //ubah data slice

	fmt.Println(slice3) //ketika di print tambah data, di sini slice buat array baru karena sudah penuh
	fmt.Println(slice2) //slice2 tetap karena pada slice3 buat array baru karena penuh
	fmt.Println(months) //array asli tidak berubah

	/*FUNCTION MAKE
	Membuat Slice Baru dengan function make
	*/
	fmt.Println("================>>> function make <<<============")
	newSlice := make([]string, 2, 5)
	newSlice[0] = "1"
	newSlice[1] = "2"
	fmt.Println(newSlice)

	/*FUNCTION COPY */
	fmt.Println("================>>> function copy <<<============")
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

}
