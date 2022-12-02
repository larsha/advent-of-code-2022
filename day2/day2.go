package day2

import (
	"fmt"
	"os"
	"strings"
)

const (
	Rock    = 1
	Paper   = 2
	Scissor = 3
)

const (
	DrawScore = 3
	WinScore  = 6
)

func gameRounds() []string {
	fileBytes, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(strings.ReplaceAll(string(fileBytes), " ", ""), "\n")

	return sliceData
}

func play(game string) (int, int) {
	m := map[string][]int{
		"AX": {Rock, Rock},
		"BY": {Paper, Paper},
		"CZ": {Scissor, Scissor},
		"AY": {Rock, Paper},
		"BX": {Paper, Rock},
		"AZ": {Rock, Scissor},
		"CX": {Scissor, Rock},
		"BZ": {Paper, Scissor},
		"CY": {Scissor, Paper},
	}

	return calcWinner(m[game][0], m[game][1])
}

func playFixed(game string) (int, int) {
	m := map[string][]int{
		"AX": {Rock, Scissor},
		"BY": {Paper, Paper},
		"CZ": {Scissor, Rock},
		"AY": {Rock, Rock},
		"BX": {Paper, Rock},
		"AZ": {Rock, Paper},
		"CX": {Scissor, Paper},
		"BZ": {Paper, Scissor},
		"CY": {Scissor, Scissor},
	}

	return calcWinner(m[game][0], m[game][1])
}

func calcWinner(s1 int, s2 int) (int, int) {
	if s1 == Rock && s2 == Scissor || s1 == Paper && s2 == Rock || s1 == Scissor && s2 == Paper {
		s1 += WinScore
	} else if s1 == s2 {
		s1 += DrawScore
		s2 += DrawScore
	} else {
		s2 += WinScore
	}

	return s1, s2
}

func Part1() (int, int) {
	opponentScore := 0
	myScore := 0

	games := gameRounds()
	for _, game := range games {
		p1, p2 := play(game)
		opponentScore += p1
		myScore += p2
	}

	return opponentScore, myScore
}

func Part2() (int, int) {
	opponentScore := 0
	myScore := 0

	games := gameRounds()
	for _, game := range games {
		p1, p2 := playFixed(game)
		opponentScore += p1
		myScore += p2
	}

	return opponentScore, myScore
}
