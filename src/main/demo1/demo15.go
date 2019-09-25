package main

import "fmt"

func zeroval(int2 int) int {
	int2 += 1

	return int2
}

func zeroptr(iptr *int) int{
	*iptr = 0

	return *iptr+1
}

func main() {
	i := 1
	fmt.Println("initial:", i)
	i =zeroval(i)
	fmt.Println("zeroval:", i)

	i =zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
