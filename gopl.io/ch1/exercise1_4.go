package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fileCounts := make(map[string]map[string]int)
	for _, filename := range os.Args[1:] {
		counts := make(map[string]int)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exercise1_4: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
		fileCounts[filename] = counts
	}
	for f, c := range fileCounts {
		for line, n := range c {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", f, n, line)
			}
		}
	}
}
