package main

import "fmt"

func main() {
	fmt.Println(fib5(5))
	fmt.Println(fib5(10))
	fmt.Println(fib5(50))
}

func fib5(n int) int {
	if n == 0 {
		return n
	}

	last := 0 //初期値はfib(0)
	next := 1 //初期値はfib(1)
	for i := 1; i < n; i++ {
		last, next = next, last+next
	}
	return next
}
