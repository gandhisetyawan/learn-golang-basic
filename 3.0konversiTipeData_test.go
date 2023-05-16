package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestKonversiTipeData(t *testing.T) {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	var name = "gandhi"
	var e byte = name[1]
	var eString string = string(e)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(eString)
}
