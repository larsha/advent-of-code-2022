package day6

import (
	"bufio"
	"os"
	"strings"
)

func parser(n int) int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	var chars []string
	for scanner.Scan() {
		r := scanner.Text()
		chars = append(chars, r)
	}

	for i := 0; i < len(chars); i++ {
		if i < n {
			continue
		}

		var msg string
		for _, k := range chars[i-(n-1) : i+1] {
			msg += k
		}

		for j, l := range msg {
			if strings.Count(msg, string(l)) >= 2 {
				break
			}

			if j == len(msg)-1 {
				return i - j + n
			}
		}
	}

	return 0
}

func Part1() int {
	return parser(4)
}

func Part2() int {
	return parser(14)
}
