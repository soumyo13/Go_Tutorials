package main

import "fmt"

func main() {
	//initializingAnArray()
	//compositeLiteral()
	//rangeOfValues()
	//appendingTheValues()
	//makeTheValues()
	multidimensionalSlices()
}

func initializingAnArray() {
	//Initializing an array
	var x [5]int
	fmt.Println(x)
	x[2] = 2 //2 is the index
	fmt.Println(x)
}

func compositeLiteral() {
	//composite literal
	//x:= type{values}
	x := []int{4, 5, 6, 7}
	fmt.Println(x)
	fmt.Println(x[1])
}

func rangeOfValues() {
	x := []int{0, 5, 6, 7}
	fmt.Println(x)

	for i, v := range x { //i is the index and v is the values in the array
		fmt.Println(i, v)
	}
	fmt.Println("Done")

	for v := range x { //i is the index and v is the values in the array
		fmt.Println(v)
	}
	fmt.Println("Done")
	for v := range x { //i is the index and v is the values in the array
		fmt.Println(x[v])
	}
	fmt.Println("Done")

	for i, v := range x { //i is the index and v is the values in the array
		fmt.Println(i, x[v], v)
		//We are not getting because the value of x[v] in second is coming as 5 and we will get an error array bound exception
	}
	fmt.Println("Done")

}

func appendingTheValues() {
	x := []int{4, 5, 6, 7}
	fmt.Println(x)
	x = append(x, 22, 23)
	fmt.Println(x)
	y := append(x, 22, 23)
	fmt.Println(y)
	z1 := []int{1, 2, 3, 4}
	z2 := []int{5, 6, 7, 8}
	z1 = append(z1, z2...)
	fmt.Println(z1)
	z1 = append(z1[2:], z1[2:]...)
	fmt.Println(z1)
}

func makeTheValues() {
	x := make([]int, 5, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("Done")

	x[0] = 1
	x[1] = 2
	x[2] = 3
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("Done")

	x = append(x, 4, 5, 6, 7, 8)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("Done")

	x = append(x, 9, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("Done")
}

func multidimensionalSlices() {
	name1 := []string{"Ram", "Sam", "Jodu", "Madhu"}
	fmt.Println(name1)
	name1 = append(name1, "Sita")
	fmt.Println(name1)

	name2 := []string{"Gita", "Mita"}
	fmt.Println(name2)

	fmt.Println(append(name1, name2...)) //Appending in a single array

	combinedNames := [][]string{name1, name2}
	fmt.Println(combinedNames) //Appending in a multiple Array

}
