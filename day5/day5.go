package day5

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	InputFile = "input.txt"
)

func moves(t string) (int, int, int) {
	var nums []int
	splitted := strings.Split(t, " ")

	for i, v := range splitted {
		n, _ := strconv.Atoi(v)

		if i%2 == 0 {
			continue
		}

		nums = append(nums, n)
	}

	// units, from, to
	return nums[0], nums[1] - 1, nums[2] - 1
}

func message(m map[int][]string) string {
	var msg string
	keys := make([]int, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		msg += m[k][0]
	}

	return msg
}

type Stacks map[int][]string

func stacks() Stacks {
	file, _ := os.Open(InputFile)
	defer file.Close()

	stacks := make(Stacks)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()

		for i, r := range t {
			if r == 91 && i%4 == 0 {
				stack := i/4 + 1 - 1
				stacks[stack] = append(stacks[stack], string(t[i+1]))
			}
		}
	}

	return stacks
}

func Part1() string {
	stacks := stacks()

	file, _ := os.Open(InputFile)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()

		if strings.HasPrefix(t, "move") {
			units, from, to := moves(t)

			for i := 1; i <= units; i++ {
				stacks[to] = append([]string{stacks[from][0]}, stacks[to]...)
				stacks[from] = append(stacks[from][:0], stacks[from][1:]...)
			}
		}
	}

	return message(stacks)
}

func Part2() string {
	stacks := stacks()

	file, _ := os.Open(InputFile)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()

		if strings.HasPrefix(t, "move") {
			units, from, to := moves(t)

			fromCopy := make([]string, len(stacks[from]))
			copy(fromCopy, stacks[from])

			stacks[to] = append(fromCopy[:units], stacks[to]...)
			stacks[from] = append(stacks[from][:0], stacks[from][units:]...)
		}
	}

	return message(stacks)
}
