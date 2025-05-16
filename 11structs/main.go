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
	fmt.Printf("Name of admin1 is %v and age is %v.", admin1.Name, admin1.Age)
}
