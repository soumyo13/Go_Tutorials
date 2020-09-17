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
	//normalFunction()
	//parametersFunctions("Soumyo")
	// returnFunction("Hello")
	// returnValue := returnFunction("Soumyo")
	// fmt.Println(returnValue)

	//multipleReturnValue1, multipleReturnValue2 := multipleReturnFunctions("Soumyo", "Karmakar")
	//fmt.Println(multipleReturnValue1, multipleReturnValue2)

	//variadicParametersFunction(1, 2, 3, 4)

	// xi := []int{1, 2, 3, 4}
	// sliceFunction(xi...)

	// defer f1()
	// f2()
	// defer f3()

	// sa1 := secretAgent{
	// 	person: person{
	// 		firstName: "Secret1",
	// 		lastName:  "Agent1",
	// 		age:       12,
	// 	},
	// 	licenseToKill: false,
	// }
	// sa1.speak()

	//Anonymous Function - Always have to define inside a variable
	// func(x int) {
	// 	fmt.Println("This is an anonymous function. I am passing an integer.", x)
	// }(42)

	//Function Expression
	// f := func() {
	// 	fmt.Println("This is an function expression...")
	// }
	// f()

	//Return a Function
	s1 := returnAFunction()
	fmt.Println(s1)
}

func normalFunction() {
	fmt.Println("This is the normal function...")
}
func parametersFunctions(s string) {
	fmt.Println("This is the parameter function which prints : " + s)
}
func returnFunction(s string) string {
	return fmt.Sprintln("This is the return function which prints: " + s)
}
func multipleReturnFunctions(fname string, lname string) (string, bool) {
	stringValues := fmt.Sprint(fname + " " + lname + " says ")
	booleanValues := true
	return stringValues, booleanValues
}
func variadicParametersFunction(numbers ...int) {
	fmt.Println(numbers)
}
func sliceFunction(numbers ...int) {
	fmt.Println(numbers)
}

//Functions
func f1() {
	fmt.Println("F1")
}
func f2() {
	fmt.Println("F2")
}
func f3() {
	fmt.Println("F3")
}

//EndOfFunctions
func (s secretAgent) speak() {
	fmt.Println("I am ", s.firstName, s.lastName)
}

//Return a Function
func returnAFunction() string {
	s := "Hello!!!"
	return s
}
