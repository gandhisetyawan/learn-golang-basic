package learn_golang_basic

import (
	"fmt"
	"reflect"
	"testing"
)

type Sample struct {
	Name string
}

func TestPackageReflect(t *testing.T) {
	sample := Sample{"gandhi"}

	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())    // melihat ada berapa field
	fmt.Println(sampleType.Field(0).Name) // melihat name field pada index ke 1

}
