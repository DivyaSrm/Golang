package main

import (
	"fmt"
	"time"
)

func run(start, end int, name string, ch chan bool) {
	for i := start; i < end; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(name, i)
	}

	ch <- true

}
func square(a int, ch chan int) {
	result := a * a
	ch <- result
}
func cube(a int, ch chan int) {
	result := a * a * a
	ch <- result
}
func main() {
	ch := make(chan bool)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go run(2, 8, "i", ch)
	go run(6, 9, "j", ch)
	go square(10, ch1)
	go cube(10, ch2)
	//time.Sleep(500 * time.Millisecond)
	fmt.Println(<-ch1 + <-ch2)
	<-ch
}
