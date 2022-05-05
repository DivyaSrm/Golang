package main

import "fmt"

func main() {
	for i := 7; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}

	var arr [9]int

	arr[0] = 12

	fmt.Println(arr)

	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	arr2 := [...]int{1, 2, 3}

	arr2[2] = 2

	fmt.Println(arr2)

	

}
