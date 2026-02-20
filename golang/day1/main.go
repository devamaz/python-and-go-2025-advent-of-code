package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data := readFile("../../input/day1.txt")
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
