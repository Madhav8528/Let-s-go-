package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://go.dev:7500/learn?pkg=bufio&func=input"

func main() {

	fmt.Println(myurl)

	//parsing the url
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["pkg"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	//its return type is *url.URL but if you print it, string would be printed so no need to convert to string.
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "go.dev",
		Path:    "/learn",
		RawPath: "user=madhav",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
	fmt.Printf("%T", partsOfUrl)

}
