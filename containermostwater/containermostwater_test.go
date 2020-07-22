package containermostwater

import "testing"

type maxAreaTestCase struct {
	height       []int
	expectedArea int
}

type areaCalculationTestCase struct {
	h            []int
	i, j         int
	expectedArea int
}

func TestArea(t *testing.T) {
	testCases := []areaCalculationTestCase{
		{[]int{0, 2, 0, 0, 0, 4}, 1, 5, 8},
		{[]int{0, 100, 100, 0, 0, 4}, 1, 2, 100},
		{[]int{0, 5, 0, 4, 0, 4}, 1, 3, 8},
		{[]int{0, 5, 0, 4, 0, 4}, 3, 1, 8},
	}
	for i, testCase := range testCases {
		area := calculateWaterArea(testCase.h, testCase.i, testCase.j)
		if testCase.expectedArea != area {
			t.Errorf("#%v: Expected %v, got %v; Input h = %v, i = %v, j = %v", i, testCase.expectedArea, area, testCase.h, testCase.i, testCase.j)
		}
	}
}

func TestMaxArea(t *testing.T) {
	testCases := []maxAreaTestCase{
		{[]int{}, 0},
		{[]int{8}, 0},
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 7, 6, 2, 5, 4, 8, 3, 8}, 49},
		{[]int{1, 0, 0, 0, 0}, 0},
		{[]int{1, 100, 100, 1, 1, 1, 1, 1, 1}, 100},
	}
	for i, testCase := range testCases {
		area := maxArea(testCase.height)
		if testCase.expectedArea != area {
			t.Errorf("#%v: Expected %v, got %v; Input %v", i, testCase.expectedArea, area, testCase.height)
		}
	}
}
