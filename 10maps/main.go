package main

import "fmt"

func main() {

	ratings := make(map[string]int)

	ratings["momo"] = 8
	ratings["choclates"] = 10
	ratings["curries"] = 8
	ratings["noodles"] = 9

	fmt.Println(ratings)
	delete(ratings, "curries")
	fmt.Printf("Rating of noodle is: %v/10\n", ratings["noodles"])

	for key, value := range ratings {
		fmt.Printf("%v has rating: %v out of 10 \n", key, value)
	}
}
