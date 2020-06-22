package sort

import (
	"cormen/util/bounds"
)

func InsertionSort(arr []int) {
	for idx, element := range arr {
		insertPos := bounds.InsertUpper(arr[0:idx], element)
		for i := idx; i > insertPos; i-- {
			arr[i] = arr[i-1]
		}
		arr[insertPos] = element
	}
}
