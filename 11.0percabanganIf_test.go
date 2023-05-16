package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestPercabanganIf(t *testing.T) {
	name := "asdadasdas"

	if name == "gandhi" {
		fmt.Println("halo", name)
	} else if name == "setyawan" {
		fmt.Println("halo", name)
	} else {
		fmt.Println(name, "kenaln doang")
	}

	//if dengan short statemen pada go
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("name benar")
	}
}
