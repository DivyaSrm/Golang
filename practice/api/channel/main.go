package main

import (
	"fmt"
	"time"
)

func myfunc(ch chan *int) {

	fmt.Println(*<-ch)

}
func main() {

	fmt.Println("start Main method")
	a := 3
	s := &a
	ch := make(chan *int)
	go myfunc(ch)
	ch <- s
	time.Sleep(1 * time.Second)
	fmt.Println("End Main method")
}
