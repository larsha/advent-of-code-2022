package day2

import (
	"testing"
)

func TestPart1(t *testing.T) {
	opponentScore, myScore := Part1()

	opponentScoreWant := 14360
	myScoreWant := 10941

	if myScore != opponentScoreWant && myScore != myScoreWant {
		t.Fatalf(`Want: %v (opponentScore), %v (myScore) --> return value: %v, %v`, opponentScoreWant, myScoreWant, opponentScore, myScore)
	}
}
