package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type secretAgent struct {
	person        //Anonymous Field Implicitly
	licenseToKill bool
}

func main() {
	//structMethod()
	//embeddedStructMethod()
	anonymousStructMethod()
}

func structMethod() {
	p1 := person{
		firstName: "Ram",
		lastName:  "K",
	}
	p2 := person{
		firstName: "Sam",
		lastName:  "L",
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.firstName, p1.lastName)
}

func embeddedStructMethod() {
	p1 := person{
		firstName: "Ram",
		lastName:  "K",
	}

	sa1 := secretAgent{
		person: person{
			firstName: "Secret1",
			lastName:  "Agent1",
			age:       12,
		},
		licenseToKill: false,
	}
	fmt.Println(p1)
	fmt.Println(sa1)
	fmt.Println(sa1.person.firstName)
	fmt.Println(sa1.firstName)
	fmt.Println(sa1.licenseToKill)
}

func anonymousStructMethod() {
	p1 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "Ram",
		lastName:  "K1",
		age:       22,
	}
	fmt.Println(p1)
	fmt.Println(p1.lastName)
}
