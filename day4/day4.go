package day4

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func transform(s string) (int, int, int, int) {
	pairs := strings.Split(s, ",")

	var f []int
	for _, p := range pairs {
		splitted := strings.Split(p, "-")
		for _, v := range splitted {
			value, _ := strconv.Atoi(v)
			f = append(f, value)
		}
	}

	return f[0], f[1], f[2], f[3]
}

func counter() (int, int) {
	file, _ := os.Open("input.txt")
	defer file.Close()

	containsCount := 0
	overlapCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p1Start, p1End, p2Start, p2End := transform(scanner.Text())

		if p1Start >= p2Start && p1End <= p2End || p2Start >= p1Start && p2End <= p1End {
			containsCount++
		}

		if p1Start <= p2End && p1Start >= p2Start || p2Start <= p1End && p2Start >= p1Start {
			overlapCount++
		}
	}

	return containsCount, overlapCount
}

func Part1() int {
	c, _ := counter()

	return c
}

func Part2() int {
	_, o := counter()

	return o
}
