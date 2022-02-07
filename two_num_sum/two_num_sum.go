package twonumsum

import (
	"sort"
)

func BinSearch(nums *[]int, target int) int {
	low := 0
	high := len(*nums)

	for low < high {
		mid := low + (high-low)/2
		if (*nums)[mid] == target {
			return mid
		} else if (*nums)[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return -1
}

func twoSum(nums []int, target int) []int {
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	sort.Ints(nums)

	var swap int = 0
	for i := 0; i < len(nums); i++ {
		compliment := target - nums[i]
		swap, nums[i] = nums[i], swap
		if a := BinSearch(&nums, compliment); a != -1 {
			// it exists, get the index in copyNums of
			// the values nums[i] and nums[a]

			swap, nums[i] = nums[i], swap

			searchStart := nums[i]
			searchEnd := compliment

			searchIndStart := 0
			searchIndEnd := 0

			for j := 0; j < len(copyNums); j++ {
				if copyNums[j] == searchStart && searchIndEnd != j {
					searchIndStart = j
				}
				if copyNums[j] == searchEnd && searchIndStart != j {
					searchIndEnd = j
				}
			}

			if searchIndStart > searchIndEnd {
				searchIndStart, searchIndEnd = searchIndEnd, searchIndStart
			}
			return []int{searchIndStart, searchIndEnd}

		}
		swap, nums[i] = nums[i], swap
	}

	return []int{-1, -1}
}

func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}
