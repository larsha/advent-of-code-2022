package day4

import (
	"testing"
)

func TestPart1(t *testing.T) {
	n := Part1()

	want := 526

	if n != want {
		t.Fatalf(`Want: %v --> return value: %v`, want, n)
	}
}

func TestPart2(t *testing.T) {
	n := Part2()

	want := 886

	if n != want {
		t.Fatalf(`Want: %v --> return value: %v`, want, n)
	}
}
