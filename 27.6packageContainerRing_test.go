package learn_golang_basic

import (
	"container/ring"
	"fmt"
	"strconv"
	"testing"
)

func TestPackageContainerRing(t *testing.T) {
	var data *ring.Ring = ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "value :" + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
