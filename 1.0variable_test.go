package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestVariable(t *testing.T) {
	//deklarasi variabel
	var name string
	var age int

	name = "gandhi"
	age = 28
	//langsung inisialisasi nilai
	var lastname = "setyawan"
	//inisialisasi nilai ulang
	name = "iwan"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(lastname)
}
