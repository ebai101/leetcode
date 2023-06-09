/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
package problems

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	dict := [128]bool{}
	length, max := 0, 0
	for i, j := 0, 0; i < len(s); i++ {
		index := s[i]
		if dict[index] {
			for ; dict[index]; j++ {
				// from j to i
				length--
				dict[s[j]] = false
			}
		}

		dict[index] = true
		length++
		if length > max {
			max = length
		}
	}
	return max
}

// @lc code=end
