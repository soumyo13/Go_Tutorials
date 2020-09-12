package main

import "fmt"

func main() {
	x := 10 //Integer
	fmt.Println(x)
	y := `hello "Soumyo"` //String
	fmt.Println(y)
	z := true // Boolean
	fmt.Println(z)

	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Printf("%T", s)
}
