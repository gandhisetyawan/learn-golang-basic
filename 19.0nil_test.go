/**Nil
- Di golang saat membuat variabel dengan tipe data tertentu , maka secara otomatis
akan di buatkan default value nya
- Namun di golang ada data nil, yakni data kosong
- nil hayna dapat dgunakan di beberapa tipe data sbb :
	- interface
	- function
	- map
	- slice
	- pointer
	- channel
*/

package learn_golang_basic

import (
	"fmt"
	"testing"
)

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": "gandhi",
		}
	}
}

func TestNil(t *testing.T) {
	data1 := NewMap("")
	var data2 map[string]string = nil // test isi nil apa an pada map

	fmt.Println(data1)
	fmt.Println(data2)

	if data1 == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(data1)
	}
}
