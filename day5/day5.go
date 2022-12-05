package day5

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
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

func parser() (Stacks, []string) {
	file, _ := os.Open("input.txt")
	defer file.Close()

	stacks := make(Stacks)
	var instructions []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()

		if strings.HasPrefix(t, "move") {
			instructions = append(instructions, t)
		}

		for i, r := range t {
			if r == 91 && i%4 == 0 {
				stack := i/4 + 1 - 1
				stacks[stack] = append(stacks[stack], string(t[i+1]))
			}
		}
	}

	return stacks, instructions
}

func Part1() string {
	stacks, instructions := parser()

	for _, instr := range instructions {
		units, from, to := moves(instr)

		for i := 1; i <= units; i++ {
			stacks[to] = append([]string{stacks[from][0]}, stacks[to]...)
			stacks[from] = append(stacks[from][:0], stacks[from][1:]...)
		}
	}

	return message(stacks)
}

func Part2() string {
	stacks, instructions := parser()

	for _, instr := range instructions {
		units, from, to := moves(instr)

		fromCopy := make([]string, len(stacks[from]))
		copy(fromCopy, stacks[from])

		stacks[to] = append(fromCopy[:units], stacks[to]...)
		stacks[from] = append(stacks[from][:0], stacks[from][units:]...)
	}

	return message(stacks)
}
