/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if val, ok := m[v]; ok && v*2 == target {
			return []int{i, val}
		}
		if val, ok := m[target-v]; ok {
			return []int{i, val}
		}
		m[v] = i
	}

	for i, v := range nums {
		if v*2 == target {
			continue
		}
		if val, ok := m[target-v]; ok {
			return []int{i, val}
		}
	}
	return nil
}

// @lc code=end

