package main

import "fmt"

//if declare with first letter capital, the variable is taken as global scope(public)
var Unknown int = 5855

func main() {
	var email string = "iamtheone@go.co.in"
	fmt.Println(email)
	fmt.Printf("The followed data type is: %T \n", email)

	var phone int64
	fmt.Println(phone)
	fmt.Printf("The followed data type is: %T \n", phone)

	var isUser bool = true
	fmt.Println(isUser)
	fmt.Printf("The followed data type is: %T \n", isUser)

	var momo float32 = 200.51215345312315
	fmt.Println(momo)
	fmt.Printf("The followed data type is: %T \n", momo)

	//auto assign of data types by lexer
	//this type of declaration(walrus) can only be done inside a function
	variablex := 5245.854556422
	fmt.Println(variablex)
	fmt.Printf("The followed data type is: %T \n", variablex)

	fmt.Println(Unknown)
	fmt.Printf("The followed data type is: %T \n", Unknown)
}
