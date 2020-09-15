package main

import "fmt"

func main() {
	initializingAnArray1()
}

//Initializing an array
func initializingAnArray() {
	var x [5]int
	fmt.Println(x)
	x[2] = 2 //2 is the index
	fmt.Println(x)

}

func initializingAnArray1() {
	var x [5]int
	fmt.Println(x)
	x[2] = 2
	fmt.Println(x)

}
