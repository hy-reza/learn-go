package main

import "fmt"

type Person struct {
	name string
	job  string
	age  int
}

//METHOD
func (p *Person) Introduce(text string) (result string) {
	p.name = "Reza"
	result = fmt.Sprintf("%s %s i am %d Y/O", text, p.name, p.age)
	return
}

func main() {
	person := Person{
		name: "张三",
		job:  "学生",
		age:  18,
	}

	fmt.Println(person.Introduce("Hi my name is"))
	fmt.Println(person)

}
