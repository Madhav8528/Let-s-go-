package main

import (
	"fmt"
	"math/rand"
	"time"
)

// can't get the same output in go playground due to some fixed operational values on the remote server rand does'nt work properly.
func main() {

	rand.NewSource(time.Now().Unix())

	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Move 1 space or can open a block")
	case 2:
		fmt.Println("Move 2 spaces")
		fallthrough
	case 3:
		fmt.Println("Move 3 spaces")
	case 4:
		fmt.Println("Move 4 spaces")
	case 5:
		fmt.Println("Move 5 spaces")
	case 6:
		fmt.Println("Move 6 and roll back again!")
	default:
		fmt.Println("Invalid input!")
	}
}
