package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	target := 9
	expectedOutput := []int{3, 4}
	output := twoSum(input, target)
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Errorf("Expected %v, got %v", expectedOutput, output)
	}
}

func TestTwoSum2(t *testing.T) {
	input := []int{1, 4, 2, 2, 5}
	target := 4
	expectedOutput := []int{2, 3}
	output := twoSum(input, target)
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Errorf("Expected %v, got %v", expectedOutput, output)
	}
}

func TestTwoSum3(t *testing.T) {
	input := []int{1, 4, 2, 2, 5}
	target := 5
	expectedOutput := []int{0, 1}
	output := twoSum(input, target)
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Errorf("Expected %v, got %v", expectedOutput, output)
	}
}

func TestTwoSum4(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9
	expectedOutput := []int{0, 1}
	output := twoSum(input, target)
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Errorf("Expected %v, got %v", expectedOutput, output)
	}
}

func TestTwoSumNoSolution(t *testing.T) {
	input := []int{1, 4, 2, 2, 5}
	target := 20
	var expectedOutput []int
	output := twoSum(input, target)
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Errorf("Expected %v, got %v", expectedOutput, output)
	}
}
