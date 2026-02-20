package main

import (
	"strconv"
)

func Part2(s []string) int {
	length := 100
	curr := 50
	count := 0

	for _, line := range s {

		num, _ := strconv.Atoi(line[1:])
		dir := 1
		if line[0] == 'L' {
			dir = -1
		}

		count += num / length
		num %= length

		for i := 0; i < num; i++ {
			curr = (curr + dir) % length
			if curr == 0 {
				count += 1
			}
		}

	}

	return count

}
