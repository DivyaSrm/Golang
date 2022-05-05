package main

import "fmt"

func main(){
	fmt.Println("hello")
	 x,_:=add(3,2)
	fmt.Println("x",x)
}
func add(a,b int) (c,d int){
	c=a+b
	d=a-b
	return 
}