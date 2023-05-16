/*
Anonymous FUCTION
-Anonymouse function : function yng tidak memiliki nama.
ini memungkinkan kita membuat function secara langsung di variabel atau parameter sebuah function
*/
package learn_golang_basic

import (
	"fmt"
	"testing"
)

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("you are blocked ", name)
	} else {
		fmt.Println("welcome ", name)
	}
}

func TestAnonymousFunction(t *testing.T) {
	//function di simpan pada variabel
	blacklist := func(name string) bool {
		return name == "asu"
	}

	registerUser("gandhi", blacklist)
	//function langsung di jadikan argument
	registerUser("asu", func(name string) bool {
		return name == "asu"
	})

	//CEK NILAI KEMBALIAN APA HASILNYA
	fmt.Println(blacklist("asu"))
}
