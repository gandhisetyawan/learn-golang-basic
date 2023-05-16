/*
  - Package Initialization
  - saat membuat package ,kita dapat membuat function yang secara otomatis di akses ketika package di akses
  - ini cocok ketika package kita berisi function - function untuk berkomunikasi dengan database
  - untuk membuat function yang secara otomatis di akses ketika package di akses,
    kita cupuk membuat function dengan nama init
*/
package learn_golang_basic

import (
	"fmt"
	"learn_golang_basic/database"
	"testing"
)

func TestPackageInitialization(t *testing.T) {
	result := database.GetDatabase()
	fmt.Println(result)
	database.ChangeAddressToIndonesia(database.AddressPointer)
}
