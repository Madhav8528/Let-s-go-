package main

import "fmt"

func main() {

	defer fmt.Println("fun")
	defer fmt.Println("is")
	fmt.Println("Go")
	learnDefer()
}

func learnDefer() {

	for i := 10; i < 16; i++ {
		if i%2 != 0 {
			defer fmt.Println(i)
		} else {
			fmt.Println(i)
		}
	}
}
