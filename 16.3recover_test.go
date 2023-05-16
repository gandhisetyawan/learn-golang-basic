/*
Recover function
-recover : function yg bisa di gunkan menangkap data panic
- Yang seharusnya ketika terjadi panic proses terhenti, dengan recover proses tetap berjalan
*/
package learn_golang_basic

import (
	"fmt"
	"testing"
)

func endAppApp() {
	message := recover()
	if message != nil {
		fmt.Println("error dengan message:", message)
	}

	fmt.Println("end app")
}

func runAppApp(error bool) {
	defer endAppApp()
	if error {
		panic("APP ERROR NIH")
	} else {
		fmt.Println("APP JALAN")
	}
}
func TestRecover(t *testing.T) {
	runAppApp(true) //aplikasi tetap jalan dengan message jika di set error
	fmt.Println("test masih jalan ke baris berikutya")
}
