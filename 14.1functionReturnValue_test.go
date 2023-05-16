package learn_golang_basic

import (
	"fmt"
	"testing"
)

func getHello(name string) string {
	if name == "" {
		return "helo bro"
	} else {
		return "helo " + name
	}
}

func TestFuncReturnValue(t *testing.T) {
	fmt.Println(getHello("gandhi")) //panggil function dan pemberian argumen
}
