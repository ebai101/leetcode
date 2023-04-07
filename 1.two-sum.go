/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package main

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if v, found := m[target-num]; found {
			return []int{v, i}
		}
		m[num] = i
	}
	return []int{-1, -1}
}

// @lc code=end
