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
		"AX": {Rock + DrawScore, Rock + DrawScore},
		"BY": {Paper + DrawScore, Paper + DrawScore},
		"CZ": {Scissor + DrawScore, Scissor + DrawScore},
		"AY": {Rock, Paper + WinScore},
		"BX": {Paper + WinScore, Rock},
		"AZ": {Rock + WinScore, Scissor},
		"CX": {Scissor, Rock + WinScore},
		"BZ": {Paper, Scissor + WinScore},
		"CY": {Scissor + WinScore, Paper},
	}

	return m[game][0], m[game][1]
}

func playFixed(game string) (int, int) {
	m := map[string][]int{
		"AX": {Rock + WinScore, Scissor},
		"BY": {Paper + DrawScore, Paper + DrawScore},
		"CZ": {Scissor, Rock + WinScore},
		"AY": {Rock + DrawScore, Rock + DrawScore},
		"BX": {Paper + WinScore, Rock},
		"AZ": {Rock, Paper + WinScore},
		"CX": {Scissor + WinScore, Paper},
		"BZ": {Paper, Scissor + WinScore},
		"CY": {Scissor + DrawScore, Scissor + DrawScore},
	}

	return m[game][0], m[game][1]
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
