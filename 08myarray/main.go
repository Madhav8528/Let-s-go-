package main

import "fmt"

func main() {

	var arr [5]int
	arr[0] = 1
	arr[1] = 1
	arr[4] = 1
	fmt.Println(arr)

	//in GO arrays size is fixed and does'nt provide dynamic nature, the value not assigned are blank or 0.
	var carList = [4]string{"Scorpio", "Fortuner", "Gwagon"}

	fmt.Println("My car list is: ", carList)
	fmt.Println("My fleet size is: ", len(carList))
}
