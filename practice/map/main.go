package main

import "fmt"

func main() {

	student := make(map[string]int)
	student["divya"] = 10
	student["viji"] = 20
	fmt.Println(student)

	employee := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	fmt.Println(employee)
	fmt.Println(employee[1])
	value, ok := employee[1]
	fmt.Println(value, ok)
	value1, ok1 := employee[0]
	fmt.Println(value1, ok1)

	for key, value := range employee {
		fmt.Println("key", key, "value", value)
	}

	delete(employee, 1)
	fmt.Println("after deleting one", employee)

	type staff struct {
		mobile   int
		location string
	}

	staff1 := staff{
		mobile:   7889890,
		location: "chennai",
	}
	staff2 := staff{
		mobile:   9009090,
		location: "banglore",
	}

	staffs := map[string]staff{
		"divya":  staff1,
		"preeti": staff2,
	}
	fmt.Println("staffs", staffs)
	fmt.Println(len(staffs))

	copy := staffs

	staff3 := staff{
		mobile:   123455,
		location: "taramani",
	}

	copy["divya"] = staff3

	fmt.Println(staffs)
}
