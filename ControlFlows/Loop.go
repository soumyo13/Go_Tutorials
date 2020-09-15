package main

import "fmt"

func main() {
	//forsingleloop()
	//forwhileloop()
	//	forifbreakwhileloop()
	//forifbreakcontinuewhileloop()
	forswitchfallthroughwhileloop()
}

func forsingleloop() {
	// for init;condition;post{}
	for i := 0; i < 10; i++ {
		fmt.Println("Hello")
	}
}

func forwhileloop() {
	//There is no while loop. But we can run the while loop like this.
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("Done")

}

func forifbreakwhileloop() {
	//There is no while loop. But we can run the while loop like this.
	x := 1
	for {
		if x > 10 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("Done")

}

func forifbreakcontinuewhileloop() {
	//Continue Keyword
	//Even Number Print
	x := 1
	for {
		x++
		if x > 10 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
	fmt.Println("Done")
}
func forswitchfallthroughwhileloop() {
	//fallthrough Keyword
	//Switch
	//fallthrough will print even if the statement is wrong
	switch {
	case true:
		fmt.Println("Prints1")
		fallthrough
	case false:
		fmt.Println("Prints2")
		fallthrough
	case (2 != 0):
		fmt.Println("Prints3")
	case (2 != 2):
		fmt.Println("Prints4")
	case (2 == 2):
		fmt.Println("Prints5")
	}
	fmt.Println("Done")
}
