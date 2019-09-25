package main

import "fmt"

func digui(n int)  int{
	if n == 0 {
		return 1
	}
	return n * digui(n-1)
}

func main()  {
	fmt.Println(digui(5))

}
