package addtwonumbers

import (
	"testing"

	"github.com/p-dizon/LeetCode/linkedlist"
)

func TestAddTwoNumbers1(t *testing.T) {
	l1 := linkedlist.ConvertToList([]int{1, 2, 3})
	l2 := linkedlist.ConvertToList([]int{1, 2, 3})
	expectedOutput := linkedlist.ConvertToList([]int{2, 4, 6})
	sum := addTwoNumbers(l1, l2)
	if _, err := linkedlist.IsEqual(expectedOutput, sum); err != nil {
		t.Error(err)
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	l1 := linkedlist.ConvertToList([]int{9})
	l2 := linkedlist.ConvertToList([]int{9})
	expectedOutput := linkedlist.ConvertToList([]int{8, 1})
	sum := addTwoNumbers(l1, l2)
	if _, err := linkedlist.IsEqual(expectedOutput, sum); err != nil {
		t.Error(err)
	}
}

func TestAddTwoNumbers3(t *testing.T) {
	l1 := linkedlist.ConvertToList([]int{9, 8})
	l2 := linkedlist.ConvertToList([]int{9, 8})
	expectedOutput := linkedlist.ConvertToList([]int{8, 7, 1})
	sum := addTwoNumbers(l1, l2)
	if _, err := linkedlist.IsEqual(expectedOutput, sum); err != nil {
		t.Error(err)
	}
}

func TestAddTwoNumbers4(t *testing.T) {
	l1 := linkedlist.ConvertToList([]int{9, 8, 9})
	l2 := linkedlist.ConvertToList([]int{9, 8})
	expectedOutput := linkedlist.ConvertToList([]int{8, 7, 0, 1})
	sum := addTwoNumbers(l1, l2)
	if _, err := linkedlist.IsEqual(expectedOutput, sum); err != nil {
		t.Error(err)
	}
}

func TestAddTwoNumbers5(t *testing.T) {
	l1 := linkedlist.ConvertToList([]int{1, 2, 3})
	l2 := linkedlist.ConvertToList([]int{1})
	expectedOutput := linkedlist.ConvertToList([]int{2, 2, 3})
	sum := addTwoNumbers(l1, l2)
	if _, err := linkedlist.IsEqual(expectedOutput, sum); err != nil {
		t.Error(err)
	}
}

func TestAddTwoNumbers6(t *testing.T) {
	l1 := linkedlist.ConvertToList([]int{1})
	l2 := linkedlist.ConvertToList([]int{1, 2, 3})
	expectedOutput := linkedlist.ConvertToList([]int{2, 2, 3})
	sum := addTwoNumbers(l1, l2)
	if _, err := linkedlist.IsEqual(expectedOutput, sum); err != nil {
		t.Error(err)
	}
}

func TestAddTwoNumbersEmptyList(t *testing.T) {
	var expectedOutput *linkedlist.ListNode
	sum := addTwoNumbers(nil, nil)
	if _, err := linkedlist.IsEqual(expectedOutput, sum); err != nil {
		t.Error(err)
	}
}
