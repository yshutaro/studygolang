package main

import "fmt"

func main() {
	fmt.Println(fib2(5))
	fmt.Println(fib2(10))
}

func fib2(n int) int {
	if n < 2 {
		return n
	}
	return fib2(n-2) + fib2(n-1)
}
