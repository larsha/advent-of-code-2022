package day1

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortedNumbers() []int {
	fileBytes, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), "\n")

	var numbers []int
	chunk := 0
	for _, s := range sliceData {
		if s == "" {
			numbers = append(numbers, chunk)
			chunk = 0
			continue
		}

		n, err := strconv.Atoi(s)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		chunk += n
	}

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	return numbers
}

func Part1() int {
	n := sortedNumbers()

	return n[0]
}

func Part2() int {
	n := sortedNumbers()

	return n[0] + n[1] + n[2]
}
