package main

import (
	"cormen/algo/sort"
	"fmt"
)

func main() {
	var arr2 []int = []int{2, 4, 5, 1, 8, 4, 0, 10, 2, 4, 2, -7}
	sort.InsertionSort(arr2)
	fmt.Printf("%v", arr2)
}
