package main

import "fmt"

func variadictest(x int, y ...int) {
	fmt.Println(x, y)
}

func test(a int, b ...int) {

	for i, v := range b {
		if v == a {

			fmt.Println(v, "found at ", i, "in", b)
		}

	}

}
func main() {
	variadictest(1, 2, 3, 4)
	variadictest(1, 2)
	test(1, 2, 3, 4, 5, 1, 1, 1)
}
