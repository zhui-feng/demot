package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func closure() func(int) int {
	var x int
	return func(i int) int {
		x++
		return  i+x
	}
}

func main()  {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	A :=closure()
	fmt.Println("x:=",A(1))
	fmt.Println("x:=",A(2))

}
