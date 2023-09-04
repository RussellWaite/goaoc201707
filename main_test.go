package main

import (
	"testing"
)

func TestSolvePart1(t *testing.T) {
	result, err := SolvePart1("./input")
	if err != nil {
		t.Error(" %w", err)
	}
	if result != "dgoocsw" {
		t.Error("result wasn't correct")
	}
}
func TestSolvePart1RealInput(t *testing.T) {
	result, err := SolvePart1("./test")
	if err != nil {
		t.Error(" %w", err)
	}
	if result != "tknk" {
		t.Error("result wasn't correct")
	}
}

func TestSolvePart2(t *testing.T) {
	result, err := SolvePart2("./test", "tknk")
	if err != nil {
		t.Error(" %w", err)
	}
	if result != 60 {
		t.Error("Expected 60, received ", result)
	}
}
func TestSolvePart2RealInput(t *testing.T) {
	result, err := SolvePart2("./input", "dgoocsw")
	if err != nil {
		t.Error(" %w", err)
	}
	if result != 1275 {
		t.Error("Expected 1275, received ", result)
	}
}
