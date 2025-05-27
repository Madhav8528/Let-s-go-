package main

import (
	"fmt"
	"mongodbapi/router"
	"net/http"
)

func main() {

	r := router.Router()
	fmt.Println("Server is getting started")
	http.ListenAndServe(":8000", r)
	fmt.Println("Server is listening on port 8000.")
}
