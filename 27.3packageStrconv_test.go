//beberapa package strconv
// strconv.parseBool(string)		: Mengubah string ke bool
// strconv.parseFloat(string) 		: Mengubah string ke float
// strconv.parseInt(string) 		: Mengubah string ke int64
// strconv.FormatBool(bool) 		: Mengubah bool ke string
// strconv.FormatFloat(float, … ) 	: Mengubah float64 ke string
// strconv.FormatInt(int, … ) 		: Mengubah int64 ke string

package learn_golang_basic

import (
	"fmt"
	"strconv"
	"testing"
)

func TestPackageStrconv(t *testing.T) {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error:", err.Error())
	}

	var value string = strconv.FormatInt(1000000, 10)
	fmt.Println(value)
}
