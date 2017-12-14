package main

import (
	"fmt"
)

func ShowMessage (message string) {
	fmt.Println(message)
}

func ShowMessage2 (message string) {
	fmt.Println(message)
}

func main() {
	myMessage := "hello world!"

	ShowMessage(myMessage)
}
