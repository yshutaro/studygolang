package main

import "fmt"

var memo = map[int]int{
	0: 0,
	1: 1,
}

func main() {
	fmt.Println(fib3(5))
	fmt.Println(fib3(10))
	fmt.Println(fib3(50))
}

func fib3(n int) int {
	if _, ok := memo[n]; !ok {
		memo[n] = fib3(n-1) + fib3(n-2) //メモ化
	}
	return memo[n]
}
