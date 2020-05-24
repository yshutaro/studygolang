package main

import "fmt"

func main() {
	fmt.Println(calculatePi(1000000))
}

func calculatePi(nTerms int) float64 {
	numerator := 4.0
	denominator := 1.0
	operation := 1.0
	pi := 0.0
	for i := 0; i < nTerms; i++ {
		pi += operation * (numerator / denominator)
		denominator += 2.0
		operation *= -1.0
	}
	return pi
}
