package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Advent of Code 2015")
	fmt.Println("- Year: 2015")
	fmt.Println("- Day: 01")

	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(dat)

	fmt.Println("- Part: 1")
	res1 := Part1(input)
	fmt.Println(res1)

	fmt.Println("- Part: 2")
	res2 := Part2(input)
	fmt.Println(res2)
}
