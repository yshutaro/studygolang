package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {
	minus := false
	y := int(math.Abs(float64(x)))
	if y > int(math.Pow(2, 31))-1 {
		return 0
	}

	if y != x {
		minus = true
	}
	b := []byte(strconv.Itoa(y))
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	z, _ := strconv.Atoi(string(b))
	if z > int(math.Pow(2, 31))-1 {
		return 0
	}
	if minus {
		z = -1 * z
	}
	return z
}
