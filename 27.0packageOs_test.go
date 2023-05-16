package learn_golang_basic

import (
	"fmt"
	"os"
	"testing"
)

func TestPackageOs(t *testing.T) {
	args := os.Args
	fmt.Println(args)

	fmt.Println(args[1])

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}

	test_username := os.Getenv("APP_USER")
	test_password := os.Getenv("APP_PASS")
	fmt.Println(test_username)
	fmt.Println(test_password)
}
