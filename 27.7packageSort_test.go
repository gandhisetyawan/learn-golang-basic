package learn_golang_basic

import (
	"fmt"
	"sort"
	"testing"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func TestPackageSort(t *testing.T) {
	users := []User{
		{"gandhi", 29},
		{"setyawan", 30},
		{"iwan", 27},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
