package day2

import (
	"fmt"
	"os"
	"strings"
)

const (
	RockScore    = 1
	PaperScore   = 2
	ScissorScore = 3
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
		return RockScore + DrawScore, RockScore + DrawScore
	case "BY":
		return PaperScore + DrawScore, PaperScore + DrawScore
	case "CZ":
		return ScissorScore + DrawScore, ScissorScore + DrawScore
	case "AY":
		return RockScore, PaperScore + WinScore
	case "BX":
		return PaperScore + WinScore, RockScore
	case "AZ":
		return RockScore + WinScore, ScissorScore
	case "CX":
		return ScissorScore, RockScore + WinScore
	case "BZ":
		return PaperScore, ScissorScore + WinScore
	case "CY":
		return ScissorScore + WinScore, PaperScore
	}

	return 0, 0
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

func Part2() int {
	panic("not implemented yet")
}
