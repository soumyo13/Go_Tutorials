package main

import "fmt"

func main() {
	fmt.Print("Hello")            //This print will not include a new line after the execution is done.
	fmt.Println("Hello")          //This println will include a new line after the execution is done.
	fmt.Println("Hello", 2, true) //This println will include a new line after the execution is done.
}
