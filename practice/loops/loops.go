package main

import "fmt"

func main(){
	n:=5
	outer:
		for i:=0;i<n;i++ {
			
			for j:=0;j<n;j++ {
				fmt.Println(i,j)
				if i==j {
				break outer	
				}
				

			}
			fmt.Println()
		}
}