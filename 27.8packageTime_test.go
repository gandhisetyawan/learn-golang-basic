package learn_golang_basic

import (
	"fmt"
	"testing"
	"time"
)

func TestPackageTime(t *testing.T) {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2023, 3, 22, 6, 27, 28, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	time, _ := time.Parse(layout, "1994-04-11")
	fmt.Println(time)
}
