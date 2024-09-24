package countnegatives

func countNegatives(grid [][]int) int {
	result := 0

	for _, arr := range grid {
		for i := len(arr) - 1; i >= 0; i-- {
			if arr[i] < 0 {
				result++
			} else {
				break
			}
		}
	}

	return result
}
