package main

import "fmt"

func main() {

	var ptr *int
	fmt.Println(ptr)

	var str string = "Go is fun"
	var newStr string = "Go is now more fun"

	var ptr1 = &str
	var ptr2 = &newStr

	fmt.Println("Address of str is: ", ptr1)
	fmt.Println("Original str is: ", *ptr1)
	fmt.Println("Address of newStr is: ", ptr2)
	fmt.Println("Original newStr is: ", *ptr2)

	var newestStr = *ptr1 + " " + *ptr2
	var newptr = &newestStr

	fmt.Println("Address of newestStr is: ", newptr)
	fmt.Println("Original newestStr is: ", *newptr)
}
