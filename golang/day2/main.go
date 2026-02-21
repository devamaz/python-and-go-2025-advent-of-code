package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readFile("../../input/day2.txt")
	fmt.Println(Part1(data))
	fmt.Println(Part2(data))
}

func readFile(filename string) string {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		return scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return ""

}

func Part1(s string) int {

	count := 0

	for rangeStr := range strings.SplitSeq(s, ",") {

		splitRange := strings.Split(rangeStr, "-")
		start, _ := strconv.Atoi(splitRange[0])
		end, _ := strconv.Atoi(splitRange[1])

		for i := start; i <= end; i++ {
			id := strconv.Itoa(i)
			mid := len(id) / 2
			if id[mid:] == id[:mid] {
				count += i
			}
		}
	}

	return count

}

func Part2(s string) int {

	count := 0

	for rangeStr := range strings.SplitSeq(s, ",") {

		start, _ := strconv.Atoi(strings.Split(rangeStr, "-")[0])
		end, _ := strconv.Atoi(strings.Split(rangeStr, "-")[1])

		for i := start; i <= end; i++ {
			id := strconv.Itoa(i)
			doubled := id + id
			if strings.Contains(doubled[1:len(doubled)-1], id) {
				count += i
			}
		}
	}

	return count

}
