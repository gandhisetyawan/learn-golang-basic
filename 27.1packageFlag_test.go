package learn_golang_basic

import (
	"flag"
	"fmt"
	"testing"
)

func TestPackageFlag(t *testing.T) {
	host := flag.String("host", "localhost", "put your database host")
	username := flag.String("username", "root", "put your database username")
	password := flag.String("password", "root", "put your database password")
	number := flag.Int("number", 1, "put your NUMBER")

	flag.Parse()

	fmt.Println("host:", *host)
	fmt.Println("username:", *username)
	fmt.Println("password:", *password)
	fmt.Println("NUMBER:", *number)
}
