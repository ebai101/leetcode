/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */
package problems

import "fmt"

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	maxLen := max(len(nums1), len(nums2))
	// totalLen := len(nums1) + len(nums2)

	/*
		median should be:
		cA[int(totalLen/2)] for odd totalLen
		(cA[totalLen/2-1] + cA[totalLen/2]) / 2 for even totalLen
	*/
	var compArray []int

	for i, j := 0, 0; i < maxLen; {
		if nums1[i] > nums2[j] {
			compArray = append(compArray, nums1[i])
			i++
		} else {
			compArray = append(compArray, nums2[j])
			j++
		}
	}
	if len(compArray)%2 == 0 {
		return float64(compArray[len(compArray)/2-1]) + float64(compArray[len(compArray)/2])/2
	} else {
		return float64(compArray[int(len(compArray)/2)])
	}
}

// @lc code=end

func Problem4Test() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
