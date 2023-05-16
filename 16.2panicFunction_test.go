/*
PANIC FUNCTION
- function yang di gunakan untuk menghentikan program
- panic function biasanya di panggil ketika terjadi error pada saat program error
- saat panic function dipanggil, program terhenti , namun defer function tetap di eksekusi
*/
package learn_golang_basic

import (
	"fmt"
	"testing"
)

func endApp() {
	fmt.Println("end app")
}

func runApps(error bool) {
	defer endApp()
	if error {
		panic("APP ERROR")
	} else {
		fmt.Println("APP JALAN")
	}
}
func TestPanicFunction(t *testing.T) {
	runApps(true) //aplikasi terhenti karena terjadi panic
}
