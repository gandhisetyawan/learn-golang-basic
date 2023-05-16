package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestPercabanganSwitch(t *testing.T) {
	name := "iwan"

	switch name {
	case "gandhi":
		fmt.Println("halo", name)
	case "iwan":
		fmt.Println("halo", name)
	default:
		fmt.Println("kenallan doang")
	}

	//SWITCH DENGAN SHORT STATEMENT
	//format : switch short statemen; kondisi{}
	fmt.Println("------------------Switch dengan statement------------------")
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("name terlalu panjang")
	case false:
		fmt.Println("name sudah benar")
	}

	fmt.Println("------------------Switch tanpa kondisi------------------")
	//SWITCH TANPA KONDISI
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("name terlalu panjang")
	case length > 5:
		fmt.Println("name sedang")
	default:
		fmt.Println("name benar")
	}
}
