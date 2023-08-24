/*
FUNCTION MAP
Operasi						Keterangan
---------------------------------------------------------
len(map)						Mendapatkan jumlah data di map
map[key]						Mendapatkan data dengan key
map[key] = value				Mengubah atau update data map , jika key belum ada tambah jika sudah ada update
make(map[TypeKey]TypeValue)		Membuat map baru
delete(map, key)				Hapus data map dengan key
*/
package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	//MAP DENGAN DATA
	fmt.Println("===========>>> map dengan data <<<==========")
	person := map[string]string{
		"name":     "gan",
		"lastname": "set",
	}
	fmt.Println(person)
	fmt.Println(person["name"])

	//MENAMBAH DATA MAP
	fmt.Println("===========>>> tambah data map <<<==========")
	person["status"] = "belum kawin"
	fmt.Println(person)

	//MENGUBAH DATA MAP
	fmt.Println("===========>>> edit data map <<<==========")
	person["lastname"] = "setyawan"
	fmt.Println(person)
	fmt.Println("panjang map :", len(person))

	//MAP TANPA DATA SEBELUMNYA
	fmt.Println("===========>>> Map tanpa data sebelumnya <<<==========")
	book := make(map[string]string)
	//tambah data pada map
	book["title"] = "belajar golang"
	book["author"] = "iwan"
	book["error"] = "error ini"
	fmt.Println(book)

	//hapus data map
	fmt.Println("===========>>> hapus data map <<<==========")
	delete(book, "error")
	fmt.Println(book)
	fmt.Println("panjang map book :", len(book))
}
