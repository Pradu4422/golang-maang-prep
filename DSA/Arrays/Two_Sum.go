package main

import "fmt"

func two_sum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func two_sum_hash_map(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		component := target - num

		if index, found := m[component]; found {
			return []int{index, i}
		}
		m[num] = i
	}

	return nil
}

func two_sum_two_pointer(nums []int, target int) []int {

	a := 0
	b := len(nums) - 1

	for i := 0; i < len(nums); i++ {
		if nums[a]+nums[b] == target {
			return []int{a, b}
		} else if nums[a]+nums[b] > target {
			b--
		} else if nums[a]+nums[b] < target {
			a++
		}
	}

	return nil
}
func main() {
	nums := []int{2, 7, 11, 12, 13}
	target := 9
	output := two_sum(nums, target)
	output1 := two_sum_hash_map(nums, target)
	output2 := two_sum_two_pointer(nums, target)

	fmt.Println(output, output1, output2)
}
