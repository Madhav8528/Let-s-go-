package main

import "fmt"

//can't declare funtion inside function
func main() {
	start()
	result, mes := addInfinite(1, 2, 3, 4, 5)

	fmt.Println(result)
	fmt.Println(mes)
}

func start() {
	fmt.Println("Start the function.")
}

func addInfinite(val ...int) (int, string) {

	temp := 0

	for _, v := range val {
		temp += v
	}

	return temp, "Another return type"
}
