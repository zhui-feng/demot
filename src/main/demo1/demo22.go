package main

import "fmt"

func main()  {
	messages := make(chan string, 2)

	//fmt.Println(<-messages)
	//fmt.Println(<-messages)
	go func() {messages <- "channel"}()
	go func() {messages <- "buffered"}()
	messages <- "buffered--"
	//messages <- "channel--"




	fmt.Println(<-messages)
	fmt.Println(<-messages)


}
