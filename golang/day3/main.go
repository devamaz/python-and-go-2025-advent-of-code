package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	data := readFile("../../input/day3.txt")
	fmt.Println(Part1(data))
	fmt.Println(Part2(data))
}

func readFile(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return lines
}

func Part1(data []string) int {

	count := 0

	for _, line := range data {
		best, _ := strconv.Atoi(string(line[0]))
		largest := math.MinInt64

		for i := 1; i < len(line); i++ {
			current, _ := strconv.Atoi(string(line[i]))
			largest = int(math.Max(float64(largest), float64(best*10+current)))
			best = int(math.Max(float64(current), float64(best)))
		}

		count += largest

	}

	return count
}

func Part2(data []string) int {
	count := 0
	turns := 12

	for _, line := range data {
		best := make([]int, turns)
		for i := range best {
			best[i] = -1
		}

		for i := 0; i < len(line); i++ {
			current := int(line[i] - '0')
			for j := turns - 1; j >= 0; j-- {
				if j == 0 {
					best[0] = max(best[0], current)
				} else if best[j-1] != -1 {
					best[j] = max(best[j], best[j-1]*10+current)
				}
			}
		}

		count += best[turns-1]

	}
	return count
}
