/**
+Interface
-interface : tipe data abstrak, yang tidak dapat di implementasikan langsung.
-interface berisikan definisi-definisi method
biasanya interface digunakan sebagi kontrak

+ Implementasi Interface
-setiap tipe data yang sesuai dengan kontrak interface, secara otomatis
 dianggap sebagai interface tersebut
-tidak perlu implementasi interface secara manual
*/

package learn_golang_basic

import (
	"fmt"
	"testing"
)

type HasName interface {
	GetName() string
}
type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func SayHello(hasName HasName) {
	fmt.Println("helo", hasName.GetName())

}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func TestInterface(t *testing.T) {
	var person Person = Person{
		Name: "gandhi",
	}

	var animal Animal = Animal{
		Name: "kelinci",
	}

	SayHello(person)
	SayHello(animal)
}

type Induk interface {
	Cetak() string
}

type Zz struct {
	Name string
}
type Yy struct {
	Name string
}

func (zz Zz) Cetak() string {
	return zz.Name
}
func (yy Yy) Cetak() string {
	return yy.Name
}

func dataData(i Induk) {
	fmt.Println("implementasi interface ", i.Cetak())
}

func TestXxx(t *testing.T) {
	result1 := Zz{
		Name: "Test1 data lagi",
	}
	fmt.Println(result1.Name) //print data struct
	result2 := Yy{
		Name: "Test2 data lagi",
	}
	fmt.Println(result2.Name) //print data struct

	dataData(result1)
	dataData(result2)

}
