package main

import "fmt"

func main() {
	nums := []int{3, 1, 4, 7, 5, 8, 9, 2, 6}
	fmt.Println(selectionSort(nums))
}

func selectionSort(nums []int) []int {
	for i, _ := range nums {
		idx, _ := findMinimumNum(nums[i:len(nums)])
		nums[i], nums[i+idx] = nums[i+idx], nums[i]
	}
	return nums
}

func findMinimumNum(a []int) (idx, n int) {
	n = a[0]
	idx = 0
	for i, v := range a {
		if n > v {
			n = v
			idx = i
		}
	}
	return idx, n
}
