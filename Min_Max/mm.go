package main

import "fmt"

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	fmt.Println(Min(15, 20))
	fmt.Println(Max(15, 20))
}
