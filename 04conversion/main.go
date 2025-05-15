package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Hi!, we are testing inputs from keyboard, please provide an input.")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating!!", input)

	newRating, err := strconv.ParseInt(strings.TrimSpace(input), 0, 64)

	if err != nil {
		fmt.Println("Something went wrong", err)
	} else {
		fmt.Println("New rating: ", newRating+1)
	}

}
