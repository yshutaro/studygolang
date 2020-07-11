package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var sLens []int
	for _, v := range strs {
		sLens = append(sLens, len(v))
	}
	min := sLens[0]
	for _, v := range sLens {
		if min > v {
			min = v
		}
	}
	if min == 0 {
		return ""
	}
	result := ""
	for i := 0; i < min; i++ {
		flg := false
		for j := 0; j < len(strs)-1; j++ {
			if strs[j][i] == strs[j+1][i] {
				flg = true
			} else {
				flg = false
				break
			}
		}
		if flg {
			result += string(strs[0][i])
		} else {
			break
		}
	}
	return result
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
