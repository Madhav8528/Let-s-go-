package main

import (
	"fmt"
	"sync"
)

var jojo []int

func main() {

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	wg.Add(3)
	func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("One func")
		mut.Lock()
		jojo = append(jojo, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Two func")
		mut.Lock()
		jojo = append(jojo, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Third func")
		mut.Lock()
		jojo = append(jojo, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	fmt.Println(jojo)
	wg.Wait()
}
