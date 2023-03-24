package main

import "fmt"

type Person struct {
	name, job string
	age       int
}

func main() {
	person := Person{
		name: "张三",
		job:  "学生",
		age:  18,
	}

	person.name = "张三"
	person.job = "学生"
	person.age = 18

	fmt.Printf("My name is %s i'am %s %d y/o \n", person.name, person.job, person.age)

	//anonymous struck
	nissan := struct {
		model string
		year  int
	}{model: "AKB48", year: 2021}

	//slice of struck
	human := []Person{
		{name: "张三",
			job: "学生",
			age: 18,
		},
		{name: "张三",
			job: "学生",
			age: 18,
		},
	}

	//slice of anonymous stuck
	cars := []struct {
		model string
		years int
	}{
		{model: "AKB48", years: 2021},
		{model: "AKB47", years: 2021},
		{model: "AKB46", years: 2021},
	}

	_, _, _ = nissan, cars, human

}
