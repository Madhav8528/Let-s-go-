package main

import "fmt"

type Admin struct {
	Name       string
	Email      string
	IsVerified bool
	Age        int
	Status     string
}

func main() {

	admin1 := Admin{"Madhav", "madhav@mickensy.co.in", true, 19, "active"}
	fmt.Println(admin1)
	fmt.Printf("Admin1 details are: %+v\n", admin1)
	fmt.Printf("Name of admin1 is %v and age is %v.\n", admin1.Name, admin1.Age)
	admin1.checkStatus()
}

//this is a method that can append with a struct which we can use in functionalities like getter setter and defining tokens while authorization
func (a Admin) checkStatus() {
	if a.Status == "active" {
		fmt.Println("Admin is active.")
	} else {
		fmt.Println("Admin is not active.")
	}
}
