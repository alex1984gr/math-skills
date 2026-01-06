package pipeline

import "sort"

// CalculateMedian takes a slice of ints and returns the median as int
func CalculateMedian(numbers []int) int {
	if len(numbers) == 0 {
		return 0 // Return 0 if slice is empty
	}

	nums := make([]int, len(numbers))
	copy(nums, numbers) // Copy slice to avoid modifying original
	sort.Ints(nums)     // Sort numbers in ascending order

	mid := len(nums) / 2

	if len(nums)%2 == 0 {
		// If even, average of two middle numbers
		return (nums[mid-1] + nums[mid] + 1) / 2 // +1 for rounding
	}
	return nums[mid] // If odd, return middle number
}
