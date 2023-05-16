/*
*Access Modifier
-GOlang untuk menentukan Access Modifier cukup menggunakan nama function atau variabel
- Jika Nama di awali Huruh Besar , maka dapat di akses dari package lain.
- Jika nama di awali Huruf Kecil, maka tidak dapat di akses dari package lain.
*/
package learn_golang_basic

import (
	"fmt"
	"learn_golang_basic/helper"
	"testing"
)

func TestAccessModifier(t *testing.T) {
	helper.SayHello("iwan")
	// helper.sayGoodBay("gandhi") //error gak bisa di panggil
	fmt.Println(helper.LastName)
}
