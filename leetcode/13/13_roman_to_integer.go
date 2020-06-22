package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToInt("III"))
}

func romanToInt(s string) int {
	romanSymbols := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	n := 0
	for i, _ := range s {
		if i == 0 {
			n = n + romanSymbols[string(s[i])]
		} else {
			if romanSymbols[string(s[i-1])] >= romanSymbols[string(s[i])] {
				n = n + romanSymbols[string(s[i])]
			} else {
				n = n + romanSymbols[string(s[i])] - 2*romanSymbols[string(s[i-1])]
			}
		}

	}
	return n
}
