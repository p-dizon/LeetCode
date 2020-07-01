package linkedlist

import "testing"

func TestLinkedList(t *testing.T) {
	input := []int{1, 2, 3}
	list := ConvertToList(input)
	for i, node := 0, list; node != nil; node = node.Next {
		if node.Val != input[i] {
			t.Errorf("Expected %v, got %v", input[i], node.Val)
		}
		i++
	}
}

func TestEmptyLinkedList(t *testing.T) {
	var input []int
	list := ConvertToList(input)
	if list != nil {
		t.Errorf("Expected list to be nil")
	}
}

func TestIsEqual(t *testing.T) {
	l1 := ConvertToList([]int{1, 2})
	l2 := ConvertToList([]int{1, 2})
	match, err := IsEqual(l1, l2)
	if match == false {
		t.Errorf("Expected match to be true")
	}
	if err != nil {
		t.Errorf("Expected err to be nil")
	}
}

func TestIsEqualFalse1(t *testing.T) {
	l1 := ConvertToList([]int{1, 3})
	l2 := ConvertToList([]int{1, 2})
	match, err := IsEqual(l1, l2)
	if match == true {
		t.Errorf("Expected match to be false")
	}
	if err == nil {
		t.Errorf("Expected err to be set")
		return
	}
	expectedError := "Different values for index 1; l1.Val = [3], l2.Val = [2]"
	if err.Error() != expectedError {
		t.Errorf("Expected [%v], got [%v]", expectedError, err.Error())
	}
}

func TestIsEqualFalse2(t *testing.T) {
	l1 := ConvertToList([]int{1, 2, 4})
	l2 := ConvertToList([]int{1, 2})
	match, err := IsEqual(l1, l2)
	if match == true {
		t.Errorf("Expected match to be false")
	}
	if err == nil {
		t.Errorf("Expected err to be set")
	}
	expectedError := "Node count mismatch"
	if err.Error() != expectedError {
		t.Errorf("Expected [%v], got [%v]", expectedError, err.Error())
	}
}
