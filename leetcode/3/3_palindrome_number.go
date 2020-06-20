package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(1534236469))
	fmt.Println(isPalindrome(9646336469))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	org := x
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		rev = rev*10 + pop
	}
	fmt.Printf("org = %d\n", org)
	fmt.Printf("rev = %d\n", rev)
	return org == rev
}
