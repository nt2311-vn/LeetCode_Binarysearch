package countnegatives

func reverseBinarySearch(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r {
		m := l + (r-l)/2
		if arr[m] >= 0 {
			l = m + 1
		} else {
			r = m
		}
	}

	return len(arr[l:])
}

func countNegatives(grid [][]int) int {
	result := 0

	for _, arr := range grid {
		result += reverseBinarySearch(arr)
	}
	return result
}
