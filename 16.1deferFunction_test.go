/*
defer function
  - function yang dapat di jadwalkan,
    untuk dieksekusi setelah sebuah function selesai di eksekusi
  - Defer Function akan selalu dieksekusi walaupun terjadi error di function yang di eksekusi
*/
package learn_golang_basic

import (
	"fmt"
	"testing"
)

func logging() {
	fmt.Println("finish call function")
}

func runApp(value int) {
	defer logging() // fungsi defer tetap di eksekusi di akhir walau terjadi error atau tidak
	fmt.Println("run application")
	result := 10 / value
	fmt.Println(result)

}
func TestDeferFunction(t *testing.T) {
	runApp(0) //error
}
