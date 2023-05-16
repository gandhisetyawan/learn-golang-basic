package learn_golang_basic

import (
	"fmt"
	"testing"
)

func TestTypeDeclarations(t *testing.T) {
	type noKtp string

	var namberKtp noKtp = "101012154646528"

	fmt.Println(namberKtp)
}
