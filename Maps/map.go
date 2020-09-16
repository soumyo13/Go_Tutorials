package main

import "fmt"

func main() {
	//mapCheck()
	//mapAddCheck()
	mapDeleteCheck()
}

func mapCheck() {
	person := map[string]int{
		// name and age
		"Ram":   22,
		"Sam":   34,
		"Jodu":  33,
		"Madhu": 44,
	}
	fmt.Println(person)
	fmt.Println(person["Ram"])

	fruits := map[string]string{
		// fruits and taste
		"Mango":      "sweet",
		"Litchi":     "sour",
		"Watermelon": "sweet",
		"Lemon":      "sour",
	}
	fmt.Println(fruits)

	students := map[int]string{
		// roll number and student name
		1: "Student1",
		2: "Student2",
		3: "Student3",
		4: "Student4",
	}
	fmt.Println(students)
	fmt.Println(students[0]) // For key 0 there is no student name so it is printing blank

	value, ok := students[0]
	fmt.Println(value, ok)

	if value, ok := students[0]; !ok {
		fmt.Println(value, ok)
		fmt.Println("Student is not there as the key is not defined")
	}
}

func mapAddCheck() {
	person := map[string]int{
		// name and age
		"Ram":   22,
		"Sam":   34,
		"Jodu":  33,
		"Madhu": 44,
	}
	fmt.Println(person)
	fmt.Println(person["Ram"])

	person["Sita"] = 27
	fmt.Println(person)
	fmt.Println(person["Sita"])

	for k, v := range person {
		fmt.Println(k, v)
	}
}

func mapDeleteCheck() {
	person := map[string]int{
		// name and age
		"Ram":   22,
		"Sam":   34,
		"Jodu":  33,
		"Madhu": 44,
	}
	fmt.Println(person)
	fmt.Println(person["Ram"])

	delete(person, "Ram")
	fmt.Println(person)

	person["Sita"] = 27
	fmt.Println(person)
	fmt.Println(person["Sita"])

	for k, v := range person {
		fmt.Println(k, v)
	}
}
