package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestPerulangan(t *testing.T) {
	// FOR
	counter := 1

	for counter <= 3 {
		fmt.Println("perulangan ke", counter)
		counter++
	}
}

func TestPerulanganDenganStatement(t *testing.T) {
	/*FOR DENGAN STATEMENT
	  Terdapat 2 statemen dalam for
	  	- init statemen : statemen sebelum for di eksekusi
	  	- post statemen	: statemen yang selalu di eksekusi di akhir tiap perulangan
	*/
	data := []string{"agung", "setyawan", "aiwan", "awan"}
	fmt.Println(data[3])
	fmt.Println(len(data))
	for i := 0; i < len(data); i++ {
		fmt.Println("perulangan ke", i, "adalah", data[i])
	}
}

func TestForRangeInSlice(t *testing.T) {
	/*FOR RANGE PADA TIPE DATA SLICE
	  - Bisa di gunakan untuk iterasi semua data collection
	  - data collection array, slice, map
	*/
	countrys := []string{"indonesia", "london", "jepang"}
	// jika index di pakai
	for index, country := range countrys {
		fmt.Println("index", index, "=", country)
	}

	//jika index tidak di pakai bisa di ganti dengan _
	for _, country := range countrys {
		fmt.Println(country)
	}
}

func TestForRangeInMap(t *testing.T) {
	book := make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "iwan"

	for key, value := range book {
		fmt.Println(key, "=", value)
	}
}
