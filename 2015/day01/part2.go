// Year: 2015
// Day: 1
// Source: https://adventofcode.com/2015/day/1
//
//- Part Two ---
// Description:
// Now, given the same instructions, find the position of the first character that causes him to enter the basement (floor -1). The first character in the instructions has position 1, the second character has position 2, and so on.
//
// For example:
//
// ) causes him to enter the basement at character position 1.
// ()()) causes him to enter the basement at character position 5.
// What is the position of the character that causes Santa to first enter the basement?

package main

func Part2(input string) int {
	floor := 0
	base_idx := 0

	for idx, char := range input {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}

		if floor == -1 {
			base_idx = idx + 1
			break
		}
	}

	return base_idx
}
