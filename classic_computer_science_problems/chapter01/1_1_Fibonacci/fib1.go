package main

import "fmt"

func main() {
	fmt.Println(fib1(5))
}

func fib1(n int) int {
	return fib1(n-1) + fib1(n-2)
}
