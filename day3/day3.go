package day3

import (
	"os"
	"strings"
	"unicode"
)

func getRucksacks() []string {
	fileBytes, _ := os.ReadFile("input.txt")
	sliceData := strings.Split(string(fileBytes), "\n")

	return sliceData
}

func getRunePriority(c rune) int {
	if unicode.IsLower(c) {
		return int(c) - 96
	}

	return int(c) - 38
}

type Set map[int]struct{}

func priority(part string) Set {
	comp := make(Set)

	for _, p := range part {
		comp[getRunePriority(p)] = struct{}{}
	}

	return comp
}

func Part1() int {
	rucksacks := getRucksacks()
	sum := 0

	for _, r := range rucksacks {
		charCount := 0
		for range r {
			charCount++
		}

		firstPart := r[0 : charCount/2]
		lastPart := r[charCount/2 : charCount]

		prio1 := priority(firstPart)
		prio2 := priority(lastPart)

		for p1 := range prio1 {
			for p2 := range prio2 {
				if p1 == p2 {
					sum += p1
				}
			}

		}
	}

	return sum
}

func Part2() int {
	rucksacks := getRucksacks()
	chunkSize := 3
	sum := 0

	for i := range rucksacks {
		if i%chunkSize != 0 {
			continue
		}

		var priorityChunks []Set

		for _, c := range rucksacks[i : chunkSize+i] {
			priorityChunks = append(priorityChunks, priority(c))
		}

		for c := range priorityChunks[0] {
			count := 0

			for _, p := range priorityChunks {
				for n := range p {
					if c == n {
						count++
					}
				}
			}

			if count == chunkSize {
				sum += c
			}
		}

	}

	return sum
}
