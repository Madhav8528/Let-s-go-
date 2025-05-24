package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//jsonEncoder()
	decodeJson()
}

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func jsonEncoder() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	json, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", json)
}

func decodeJson() {

	upcomingData := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
        "Price": 299,
        "website": "LearnCodeOnline.in",
        "tags": ["web-dev","js"]
    }
	`)

	var decodedData course

	isValidJsonData := json.Valid(upcomingData)
	if isValidJsonData {
		fmt.Println("Data is valid")
		json.Unmarshal(upcomingData, &decodedData)
		fmt.Printf("%#v\n", decodedData)
	} else {
		fmt.Println("Data is invalid.")
	}

	var decodedData2 map[string]interface{}

	json.Unmarshal(upcomingData, &decodedData2)
	fmt.Printf("%#v\n", decodedData2)

}
