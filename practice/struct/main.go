package main

import "fmt"

type Student struct {
	id    int
	name  string
	grade string
	Mark
}

type Mark struct {
	tamil   int
	english int
	maths   int
}

func main() {
	stu1 := Student{
		id:    1,
		name:  "divya",
		grade: "A",
		Mark: Mark{
			tamil:   12,
			english: 23,
			maths:   34},
	}
	fmt.Println(stu1)

}
