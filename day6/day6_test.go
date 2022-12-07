package day6

import (
	"testing"
)

func TestPart1(t *testing.T) {
	n := Part1()

	want := 1892

	if n != want {
		t.Fatalf(`Want: %v --> return value: %v`, want, n)
	}
}

func TestPart2(t *testing.T) {
	n := Part2()

	want := 2313

	if n != want {
		t.Fatalf(`Want: %v --> return value: %v`, want, n)
	}
}
