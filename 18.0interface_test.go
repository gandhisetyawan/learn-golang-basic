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
	var data1 Person = Person{
		Name: "gandhi",
	}

	var data2 Animal = Animal{
		Name: "kelinci",
	}

	SayHello(data1)
	SayHello(data2)
}
