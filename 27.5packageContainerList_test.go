package learn_golang_basic

import (
	"container/list"
	"fmt"
	"testing"
)

func TestPackageContainerList(t *testing.T) {
	data := list.New()
	data.PushBack("gandhi")
	data.PushBack("setyawan")
	data.PushBack("iwan")

	// fmt.Println(data.Front().Next().Value)
	// fmt.Println(data.Back().Value)

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
