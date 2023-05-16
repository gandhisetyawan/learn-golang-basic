package database

import "fmt"

var connection string

type Address struct {
	City, Province, Country string
}

func init() {
	connection = "MySQL"
	fmt.Println("func init pertama kali dipanggil ketika sebuah package di akses")
}

func GetDatabase() string {
	return connection
}

func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia" // mengubah data struct
	fmt.Println(address)
}

var address Address = Address{
	City:     "Cikarang",
	Province: "Jawa Barat",
	Country:  "",
}

var AddressPointer *Address = &address // membuat pointer
