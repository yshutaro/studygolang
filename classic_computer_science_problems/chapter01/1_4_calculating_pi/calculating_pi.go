package main

import "fmt"

func main() {
	fmt.Println(calculate_pi(1000000))
}

func calculate_pi(n_terms int) float64 {
	numerator := 4.0
	denominator := 1.0
	operation := 1.0
	pi := 0.0
	for i := 0; i < n_terms; i++ {
		pi += operation * (numerator / denominator)
		denominator += 2.0
		operation *= -1.0
	}
	return pi
}
