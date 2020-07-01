package twosum

import "fmt"

func twoSum(nums []int, target int) []int {
	tracker := make(map[int]int)
	for i, num := range nums {
		if j, keyExists := tracker[target-num]; keyExists == true {
			return []int{j, i}
		}
		tracker[num] = i
	}
	return nil
}

func Test() {
	fmt.Println("Hello from twosum")
}
