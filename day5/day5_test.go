package day5

import (
	"testing"
)

func TestPart1(t *testing.T) {
	n := Part1()

	want := "TGWSMRBPN"

	if n != want {
		t.Fatalf(`Want: %v --> return value: %v`, want, n)
	}
}

func TestPart2(t *testing.T) {
	n := Part2()

	want := "TZLTLWRNF"

	if n != want {
		t.Fatalf(`Want: %v --> return value: %v`, want, n)
	}
}
