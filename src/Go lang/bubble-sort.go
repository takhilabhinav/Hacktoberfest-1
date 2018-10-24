package main

import "fmt"

func main() {

	nums := []int{1, 9, 4, 6, 3, 8, 2, 5, 7}
	fmt.Println(BubbleSort(nums))
}

func BubbleSort(a []int) []int {
	for {
		isSorted := false
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				isSorted = true
			}
		}
		if !isSorted {
			break
		}
	}
	return a
}
