package bounds

func InsertUpper(arr []int, element int) int {
	left, right := 0, len(arr)

	for left != right {
		mid := (left + right) / 2
		if element < arr[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
