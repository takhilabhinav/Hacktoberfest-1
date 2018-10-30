package main

import "fmt"

func main() {
	nums := []int{1, 9, 4, 6, 3, 8, 2, 5, 7}
	fmt.Println(insertionSort(nums))
}

func insertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i-j-1] > a[i-j] {
				a[i-j-1], a[i-j] = a[i-j], a[i-j-1]
			} else {
				break
			}
		}
	}
	return a
}
