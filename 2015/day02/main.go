package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Present struct {
	l int
	w int
	h int
}

func convertToInt(value string) int {
	digit, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return digit
}

func main() {
	fmt.Println("Advent of Code 2015")
	fmt.Println("- Year: 2015")
	fmt.Println("- Day: 02")

	data, err := os.ReadFile("input2.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	var presents []Present
	for _, line := range lines {
		if line != "" {
			dimensions := strings.Split(line, "x")
			present := Present{
				l: convertToInt(dimensions[0]),
				w: convertToInt(dimensions[1]),
				h: convertToInt(dimensions[2]),
			}
			presents = append(presents, present)
		}
	}

	fmt.Println("- Part: 1")
	res1 := Part1(presents)
	fmt.Println(res1)

	fmt.Println("- Part: 2")
	res2 := Part2(presents)
	fmt.Println(res2)
}
