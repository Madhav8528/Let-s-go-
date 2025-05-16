package main

import "fmt"

func main() {

	bC := 1

	if bC > 1 {
		fmt.Println("You are such a dashdash")
	} else if bC < 0 {
		fmt.Println("Move out bruhh..")
	} else {
		fmt.Println("Doing good!!")
	}

	if userCount := 5; userCount > 2 {
		fmt.Println("Often user")
	}
}
