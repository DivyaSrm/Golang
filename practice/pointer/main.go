package main

import "fmt"

func main() {
	a := new(int)
	fmt.Printf("a value %d type  %T address  %v", *a, a, a)
}
