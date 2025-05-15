package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hi!, we are testing inputs from keyboard, please provide an input.")
	reader := bufio.NewReader(os.Stdin)

	//input, _ := reader.ReadBytes('\n')
	input, _ := reader.ReadString('\n')
	//input, err := reader.ReadString('\n')
	//_, err := reader.ReadString('\n')
	fmt.Println("Thanks for the rating!!", input)
	fmt.Printf("Type of rating is: %T", input)
}
