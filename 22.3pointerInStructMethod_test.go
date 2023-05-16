/*
  - Pointer di Struct Method
  - meskipun method menempel pada struct , sebenarnya data struct yang di akses di dalam method
    adalah pass by value
  - sangat di rekomendasikan menggunaakn pointer di method,
    sehingga tidak boros memory karena harus selalu duplikasi ketia memanggil method
*/
package learn_golang_basic

import (
	"fmt"
	"testing"
)

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func TestPointerStructMethod(t *testing.T) {
	var data Man = Man{
		Name: "gandhi",
	}
	data.Married()

	fmt.Println(data.Name)
}
