package main

import "fmt"

var y = "Soumyo"

var z1 int //var i int = 0
var z2 string
var z3 bool

// Declare there is a variable with the identifier "xyz" and that the variable with the identfuer "xyz" is of type "bool/string/int"
// And assigns the value of the default

//x := 100
func main() {
	var x = 10 //Integer
	fmt.Println(x)

	//y = 20 // Static Programming Language

	// The var keyword can be used outside the main function but in case of the short decalaration we cant use outside the main.
	//The scope of the var is all through the program.
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z1)
	fmt.Println(z2)
	fmt.Println(z3)
}
