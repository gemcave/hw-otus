package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	testString := "Hello, OTUS!"

	fmt.Print(reverse.String(testString))
}
