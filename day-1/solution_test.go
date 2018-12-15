package main

import "testing"

func TestGetSum(t *testing.T) {
	input, expected := []int{3, 3, 4, -2, -4}, 4
	result := getSum(input)
	if result != expected {
		t.Errorf("expected %d, but getting %v", expected, result)
	}
}

func TestGetDupF(t *testing.T) {
	input, expected := []int{3, 3, 4, -2, -4}, 10
	result := getDupF(input)
	if result != expected {
		t.Errorf("expected %d, but getting %v", expected, result)
	}
}
