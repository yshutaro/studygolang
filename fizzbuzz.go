package main

import (
	"flag"
	"fmt"
	"strconv"

	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	//fmt.Println(args)
	if len(args) != 1 {
		os.Exit(1)
	}
	s, err := strconv.Atoi(args[0])
	if err != nil {
		os.Exit(1)
	}
	if s%15 == 0 {
		fmt.Println("FizzBuzz")
	} else if s%3 == 0 {
		fmt.Println("Fizz")
	} else if s%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(s)
	}
}
