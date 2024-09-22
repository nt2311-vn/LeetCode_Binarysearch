package search_insert

func searchInsert(nums []int, target int) int {
	i := 0
	j := len(nums)

	for {
		n := (j + i) / 2

		if nums[n] == target {
			return n
		} else if nums[n] > target {
			j = n
		} else if nums[n] < target {
			i = n
		}

		if j-i <= 1 {
			if nums[i] >= target {
				return i
			} else {
				return j
			}
		}
	}
}
