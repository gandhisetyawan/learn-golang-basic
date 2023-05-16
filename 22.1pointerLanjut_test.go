/*
*
pointer lanjutan
pointer dengan operator *
- jika ingin mengubah seluruh variabel yang mengacu ke data pointer tersebut maka pakai operator *

finction new
-sebelumnya membuat pointer menggunakn opertor &
-go-lang memiliki function new yang dapat digunakan untuk membuat pointer
-namun function new hanya mengembalikan pointer data kosong , artinya tidak ada data awal
*/
package learn_golang_basic

import (
	"fmt"
	"testing"
)

type DataAddress struct {
	City, Province, Country string
}

func TestPointerLanjut(t *testing.T) {
	var address1 DataAddress = DataAddress{
		City:     "Bekasi",
		Province: "jawa Barat",
		Country:  "indonesia",
	}
	var address2 *DataAddress = &address1

	*address2 = DataAddress{
		City:     "Tulungagung",
		Province: "jawa Timur",
		Country:  "indonesia",
	}
	var address3 *DataAddress = &address1
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

}

func TestPointerWithFunctionNew(t *testing.T) {
	//membuat pointer dengan function new
	var address4 *DataAddress = new(DataAddress)
	//beri nilai atau edit pointer
	address4 = &DataAddress{
		City:     "kediri",
		Province: "jawa Timur",
		Country:  "indonesia",
	}
	*address4 = DataAddress{
		City:     "blitar",
		Province: "jawa Timur",
		Country:  "indonesia",
	}
	fmt.Println(address4)

	var addressNew *DataAddress = &DataAddress{
		City:     "jinggong",
		Province: "tai cin pin",
		Country:  "china",
	}
	var addressNew1 = addressNew
	*addressNew = DataAddress{
		City:     "patok",
		Province: "hiho",
		Country:  "china",
	}

	fmt.Println(addressNew)
	fmt.Println(addressNew1)
}
