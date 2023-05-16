package learn_golang_basic

import (
	"fmt"
	"testing"
)

func getSay(name string) string {
	return "hello " + name
}

func TestFunctionAsValue(t *testing.T) {
	say := getSay
	fmt.Println(say("gandhi"))
}
