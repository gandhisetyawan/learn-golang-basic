package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestOperasiPerbandingan(t *testing.T) {
	/*Operasi Perbandingan
	  Operator	Keterangan
	  ===============================
	  >			Lebih Dari
	  <			Kurang dari
	  >=			Lebih Besar Sama Dengan
	  <=			Lebih Kecil Sama Dengan
	  ==			Sama Dengan
	  !=			Tidak Sama dengan
	*/
	var data1 = 1
	var data2 = 1
	var result = data1 == data2

	fmt.Println(result)
}
