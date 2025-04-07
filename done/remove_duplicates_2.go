package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			nums[k] = nums[i]
			k++
		}
	}

	for i := 0; i < k; i++ {
		fmt.Printf("%d ", nums[i])
	}
	fmt.Println()

	return k
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	removeDuplicates(nums)

}
