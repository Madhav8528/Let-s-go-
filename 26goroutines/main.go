package main

import (
	"fmt"
	"net/http"
	"sync"
)

func print(s string) {

	defer wg.Done()
	res, err := http.Get(s)
	if err != nil {
		panic(err)
	}
	mut.Lock()
	test = append(test, s)
	mut.Unlock()
	fmt.Printf("%d is content length for %s\n", res.StatusCode, s)
}

// this is a type pointer
var wg sync.WaitGroup
var mut sync.Mutex

var test = []string{}

func main() {

	websites := []string{
		"https://pkg.go.dev",
		"https://onlyfans.com",
	}

	for _, web := range websites {
		go print(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(test)
}
