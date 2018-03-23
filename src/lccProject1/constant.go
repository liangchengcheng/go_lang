package main

import "fmt"

func main()  {
	const LENGHT int= 10
	const WIDTH  int= 5

	var area int
	const a, b, c  = 1, false, "string"
	area = LENGHT * WIDTH

	fmt.Println("面试为 ： %d", area)
	println()
	println(a, b, c)
}

const (
	UNKOWN = 0
	FEMALE = 1
	MALE = 2
)


