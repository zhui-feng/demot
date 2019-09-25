package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	//go func() {
	//j,more := <-jobs
	//fmt.Println("received job", j,more)
	//}()


	go func() {
		fmt.Println(time.Now(),"zz")

		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	time.Sleep(time.Second * 5)

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}



	close(jobs)

	fmt.Println("sent all jobs")

	fmt.Println(time.Now(),"xx")

	<-done
}