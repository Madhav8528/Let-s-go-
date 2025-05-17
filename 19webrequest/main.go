package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://go.dev/"

func main() {

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	dataByte, _ := io.ReadAll(res.Body)
	fmt.Println(string(dataByte))

	//every time you have to explicitly close the response.
	res.Body.Close()
}
