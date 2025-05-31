package main

import (
	"fmt"
	"log"
	"mongodbapi/router"
	"net/http"
)

func main() {

	r := router.Router()
	fmt.Println("Server is getting started")
	log.Fatal(http.ListenAndServe(":6000", r))
	fmt.Println("Server is listening on port 6000.")
}
