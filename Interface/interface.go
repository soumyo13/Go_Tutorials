package main

import "fmt"

type human interface {
	speak()
	see()
}

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println("I am in the Speak method. My name is : ", p.firstName, p.lastName)
}
func (p person) see() {
	fmt.Println("I am in the See method. My name is : ", p.firstName, p.lastName)
}

func (s secretAgent) speak() {
	fmt.Println("I am in the Speak method. My name is : ", s.firstName, s.lastName)
}
func (s secretAgent) see() {
	fmt.Println("I am in the See method. My name is : ", s.firstName, s.lastName)
}

func checkHumanity(h human) {
	fmt.Println("Calling the checkHumanity....")

	fmt.Println("Without the switch statement...")
	fmt.Println("Defining the h without conversion: ", h)

	switch h.(type) {
	case person:
		fmt.Println("Defining the h in Person after conversion: ", h.(person).firstName)
	case secretAgent:
		fmt.Println("Defining the h in secretAgent after conversion: ", h.(secretAgent).firstName)
	}
}

//func (r receiver) identifier(parameters) (return(s)) {code}
func main() {
	p1 := person{
		firstName: "Ram",
		lastName:  "K",
	}

	sa1 := secretAgent{
		person: person{
			firstName: "Secret1",
			lastName:  "Agent1",
		},
		licenseToKill: false,
	}
	//fmt.Println("Defines the p1: ", p1)
	//fmt.Println("Defines the sa1: ", sa1)

	// p1.speak()
	// p1.see()
	// sa1.speak()
	// sa1.see()

	checkHumanity(p1)
	checkHumanity(sa1)
}
