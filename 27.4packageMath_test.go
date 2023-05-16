// //beberapa package math
// math.Round(float64) 	: Membulatkan float64 keatas atau kebawah, sesuai dengan yang paling dekat
// math.Floor(float64) 	: Membulatkan float64 kebawah
// math.Ceil(float64) 	: Membulatkan float64 keatas
// math.Max(float64, float64) 	: Mengembalikan nilai float64 paling besar
// math.Min(float64, float64) 	: Mengembalikan nilai float64 paling kecil
package learn_golang_basic

import (
	"fmt"
	"math"
	"testing"
)

func TestPackageMath(t *testing.T) {
	fmt.Println(math.Floor(1.9))
	fmt.Println(math.Ceil(1.1))
	fmt.Println(math.Max(10, 50))
	fmt.Println(math.Min(2, 5))
}
