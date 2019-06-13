package main

import (
	"fmt"

	"github.com/Omar-Khawaja/circle-ci-testing/hello"
)

func main() {
	greeting := hello.Hello("Omar")
	fmt.Println(greeting)
}