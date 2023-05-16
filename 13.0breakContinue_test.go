package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestBreak(t *testing.T) {
	//break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // di gunakan untuk menghentika seluruh perulangan.
		}
		fmt.Println(i)
	}
}

func TestContinue(t *testing.T) {
	//continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // menghentikan/melewatkan perulangan dan lanjut ke perulangan selanjutnya
		}
		fmt.Println(i)
	}
}
