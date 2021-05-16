/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 * Time: O(n)
 * Space: O(min(m, n))
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	i, j, longest := 0, 0, 0
	window := make(map[byte]int)
	for j = 0; j < len(s); j++ {
		if ind, ok := window[s[j]]; ok && ind >= i {
			if longest < j-i {
				longest = j - i
			}
			i = ind + 1
		}
		window[s[j]] = j
	}
	if longest < j-i {
		longest = j - i
	}
	return longest
}

// @lc code=end

