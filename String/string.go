package main

import "fmt"

func main() {
	s := `hello "Soumyo"` //String
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	//String to byte conversion
	byte1 := []byte(s)
	fmt.Println(byte1)
	fmt.Printf("%T\n", byte1)

	//Hexa
	for i := 0; i < len(s); i++ {
		fmt.Printf("%U ", s[i])
	}
}
