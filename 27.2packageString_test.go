// Beberapa package String
// strings.Trim(string, cutset) 		: Memotong cutset di awal dan akhir string
// strings.ToLower(string) 				: Membuat semua karakter string menjadi lower case
// strings.ToUpper(string) 				: Membuat semua karakter string menjadi upper case
// strings.Split(string, separator) 	: Memotong string berdasarkan separator
// strings.Contains(string, search) 	: Mengecek apakah string mengandung string lain
// strings.ReplaceAll(string, from, to) : Mengubah semua string dari from ke to
package learn_golang_basic

import (
	"fmt"
	"strings"
	"testing"
)

func TestPackageString(t *testing.T) {
	fmt.Println(strings.Contains("gandhi setyawan", "gandhi"))
	fmt.Println(strings.Split("gandhi setyawan", " "))
	fmt.Println(strings.ToUpper("gandhi setyawan"))
	fmt.Println(strings.Trim(" gandhi setyawan ", ""))
}
