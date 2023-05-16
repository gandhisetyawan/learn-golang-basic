package learn_golang_basic

import (
	"fmt"
	"reflect"
	"testing"
)

type Samples struct {
	Name string `required:"true" max:"10"`
}

type ContohLagi struct {
	Name        string `required:"true"`
	Description string `required:"true"`
}

func isValid(obj interface{}) bool {
	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(obj).Field(i).Interface() //mengambil nilai obj/ data struct
			if value == "" {
				return false
			}
		}
	}
	return true
}

func TestPackageReflectTag(t *testing.T) {
	sample := Samples{"gandhi"}
	var contohLagi ContohLagi = ContohLagi{"gandhi", "ok nih"}
	sampleType := reflect.TypeOf(sample) //data reflect

	structField := sampleType.Field(0).Name
	structTag := sampleType.Field(0).Tag.Get("required")

	fmt.Println(structField)
	fmt.Println(structTag)

	sample.Name = "" //mengbuah nilai obj / data struct menjadi kosong
	fmt.Println(isValid(sample))

	fmt.Println(isValid(contohLagi))
}
