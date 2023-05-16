package learn_golang_basic

import (
	"fmt"
	"testing"
)

type DataCustomer struct {
	Name, Address string
	Age           int
	meried        bool
}

func TestStructLiteral(t *testing.T) {
	// var data DataCustomer = DataCustomer{
	// 	Name:    "gandhi",
	// 	Address: "indonesia",
	// 	Age:     28,
	// }
	data := DataCustomer{
		Name:    "gandhi",
		Address: "indonesia",
		Age:     29,
	}
	fmt.Println(data)
	fmt.Println(data.Name)

	//tidak di sarankan pakai cara ini, karena pemberian nilai harus urut dengan posisi di struct nya
	data1 := DataCustomer{"rinandi", "indonesia", 31, true}
	fmt.Println(data1)
}
