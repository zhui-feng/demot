package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "uu" }()

	msg := <-messages


	go func() { messages <- "ping" }()


	msg2 := <-messages

	fmt.Println(msg2)
	fmt.Println(msg)

}
