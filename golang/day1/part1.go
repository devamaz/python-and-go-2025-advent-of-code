package main

import (
	"strconv"
	"strings"
)

func Part1(s []string) int {

	length := 100
	curr := 50
	count := 0

	for _, line := range s {

		num, _ := strconv.Atoi(line[1:])

		if strings.HasPrefix(line, "L") {
			curr = (curr - num) % length

		} else {
			curr = (curr + num) % length
		}

		if curr == 0 {
			count++
		}
	}

	return count

}
