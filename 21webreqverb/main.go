package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//getRequest()
	//postJsonRequest()
	postFormRequest()
}

func getRequest() {

	const url = "http://localhost:5000/get"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	content, _ := io.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.ContentLength)
	fmt.Println(string(content))

	//alternate method
	var str strings.Builder
	byteCon, _ := str.Write(content)

	fmt.Println(byteCon)
	fmt.Println(str.String())

}

func postJsonRequest() {

	const url = "http://localhost:5000/post"

	body := strings.NewReader(`{
		"dskd" : "bdjkdbfk",
		"bsdcjks" : "ndsjkfn"
	 }`)

	response, err := http.Post(url, "application/json", body)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(response)
	fmt.Println(string(content))
}

func postFormRequest() {

	const uurl = "http://localhost:5000/post-form"

	//accepts only key value pairs
	var data = url.Values{}
	data.Add("nfdjf", "jkdn")
	data.Add("ewg", "fewf")

	resp, err := http.PostForm(uurl, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Body)
	con, _ := io.ReadAll(resp.Body)

	fmt.Println(con)
	fmt.Println(string(con))

}
