package main

import "fmt"

func main() {
	// var person map[string]string
	// person = map[string]string{"name": "Handy", "age": "19"}

	// var person map[string]string = map[string]string{"name": "Handy", "age": "19"}

	// person :=  (map[string]string)
	// person["name"] = "Handy"
	// person["hobby"] = "coding"

	// var person = map[string]string{
	// 	"name":  "Handy Reza Alfanda",
	// 	"hobby": "coding",
	// }

	person := map[string]string{"name": "Handy Reza Alfanda", "hobby": "coding", "age": "22"}

	delete(person, "age")

	for i, v := range person {
		fmt.Printf("Key %s : %v \n", i, v)
	}

	value, axist := person["age"]

	fmt.Println(value, axist)

	//comibine slice & map
	people := []map[string]string{
		{"name": "Handy", "hobby": "coding"},
		{"name": "Reza", "hobby": "coding"},
		{"name": "Alfanda", "hobby": "coding"},
	}

	_ = people

}
