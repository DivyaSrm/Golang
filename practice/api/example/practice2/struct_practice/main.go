package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var row int
var student [12][12]string

func addStudent() {
	ques := [12]string{"rollno", "name", "subject1", "subject2", "subject3", "subject4", "subject5"}

	for i := 0; i < 7; i++ {
		if i < 7 {
			fmt.Println("Enter", ques[i])
			fmt.Scanln(&student[row][i])
		}
		if i == 0 {
			rno, err := strconv.Atoi(student[row][i])
			if err != nil {
				fmt.Println("Register number must be integer")
				i--
			} else if rno == 0 {
				fmt.Println("regno can not be zero")
				i--
			} else if rno < 0 {
				fmt.Println("regno can not be negative")
				i--
			}

		}
		if i == 1 {
			rex := regexp.MustCompile("[:alpha:]")
			nam := rex.MatchString(student[row][i])
			fmt.Println(nam)
			if !nam {

				fmt.Println("Name must be alphabetics")
				i--
			}
		}
		if i >= 2 {
			s, _ := strconv.Atoi(student[row][i])
			if s > 100 {
				fmt.Println("Mark can not be more than 100")
				i--
			} else if s < 0 {
				fmt.Println("Mark can not be less than 0")
			} else {
				if s < 50 {
					a, _ := strconv.Atoi(student[row][10])
					student[row][10] = strconv.Itoa(a + 1)
					student[row][11] = student[row][11] + "," + ques[i]
				}
				c, _ := strconv.Atoi(student[row][7])
				d, _ := strconv.Atoi(student[row][i])
				student[row][7] = strconv.Itoa(c + d)

			}

		}

	}
	g, _ := strconv.Atoi(student[row][7])
	student[row][8] = strconv.Itoa(g / 5)
	if g > 90 {
		student[row][9] = "O Grade"
	} else if g > 80 {
		student[row][9] = "A Grade"
	} else if g > 70 {
		student[row][9] = "B Grade"
	} else if g > 60 {
		student[row][9] = "C Grade"
	} else if g >= 50 {
		student[row][9] = "D Grade"
	} else {
		student[row][9] = "E Grade"
	}

	row++

	fmt.Println("Student added successfully")
}

func viewAllStudent() {

	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			if student[i][j] != "" {
				fmt.Print(student[i][j], "   ")

			}

		}
		fmt.Println()
	}
}

func Search() {
	var regno int
	fmt.Println("Enter Regno")
	fmt.Scanln(&regno)
	r := strconv.Itoa(regno)
	for i := 0; i < 2; i++ {
		for j := 0; j < 12; j++ {
			if student[i][j] == r {
				fmt.Println(student[i])
				break
			}
		}

	}

}

func main() {
	var choice string
	s := true
	for s == true {
		fmt.Println("Enter your choice!!! \n 1.Add Student  2.View all student 3.Search 4.Exit")
		fmt.Scanln(&choice)

		if userselection, err := strconv.Atoi(choice); err == nil {
			switch userselection {
			case 1:
				fmt.Println("Add Student")
				addStudent()
			case 2:
				fmt.Println("View all student")
				viewAllStudent()
			case 3:
				fmt.Println("Search")
				Search()
			case 4:
				fmt.Println("Exit")
				s = false
			default:
				fmt.Println("Incorrect choice")
			}
		}
	}

}
