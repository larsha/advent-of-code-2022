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
	switch game {
	case "AX":
		return Rock, Rock
	case "BY":
		return Paper, Paper
	case "CZ":
		return Scissor, Scissor
	case "AY":
		return Rock, Paper
	case "BX":
		return Paper, Rock
	case "AZ":
		return Rock, Scissor
	case "CX":
		return Scissor, Rock
	case "BZ":
		return Paper, Scissor
	case "CY":
		return Scissor, Paper
	}

	return 0, 0
}

func playFixed(game string) (int, int) {
	switch game {
	case "AX":
		return Rock, Scissor
	case "BY":
		return Paper, Paper
	case "CZ":
		return Scissor, Rock
	case "AY":
		return Rock, Rock
	case "BX":
		return Paper, Rock
	case "AZ":
		return Rock, Paper
	case "CX":
		return Scissor, Paper
	case "BZ":
		return Paper, Scissor
	case "CY":
		return Scissor, Scissor
	}

	return 0, 0
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
		p1, p2 := calcWinner(play(game))
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
		p1, p2 := calcWinner(playFixed(game))
		opponentScore += p1
		myScore += p2
	}

	return opponentScore, myScore
}
