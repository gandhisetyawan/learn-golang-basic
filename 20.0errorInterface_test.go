/*
*Error Interface
Golang memiliki interface yang digunakan sebagai kontrak untuk membuat error.
nama interface nya error
*/
package learn_golang_basic

import (
	"errors"
	"fmt"
	"testing"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan nol")
	} else {
		return nilai / pembagi, nil

	}
}

func TestErrorInterface(t *testing.T) {
	hasil, err := Pembagian(10, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error :", err.Error())
	}
}
