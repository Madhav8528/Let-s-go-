package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Create("./newFile.txt")
	if err != nil {
		panic(err)
	}

	content := "This content is to added."

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println(length)

	data, _ := os.ReadFile(file.Name())
	fmt.Println(string(data))
	file.Close()
}
