package main

import "fmt"

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.


Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 103
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.
*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if num, ok := m[complement]; ok && num != i {
			return []int{i, m[complement]}
		}
		m[nums[i]] = i
	}

	return nil
}

func printNums(nums []int, target int) {
	fmt.Printf("nums = %v, target = %d, ans = %v\n", nums, target, twoSum(nums, target))
}

func main() {
	printNums([]int{2, 7, 11, 15}, 9)
	printNums([]int{3, 2, 4}, 6)
	printNums([]int{3, 3}, 6)
}
