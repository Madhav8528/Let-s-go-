package main

import (
	"fmt"
	"time"
)

func main() {

	currtTime := time.Now()
	fmt.Println(currtTime)

	fmt.Println(currtTime.Format("01-02-2006 15:04:05 Monday "))

	createdDate := time.Date(2004, time.July, 20, 04, 52, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
