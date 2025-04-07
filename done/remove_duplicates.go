package main

// func removeDuplicates(nums []int) int {
// 	ns := len(nums)
// 	cnt := 0
// 	i := 0
// 	for i < ns-1 {
// 		if nums[i+1] == nums[i] {
// 			nums[i] = -111
// 			cnt++
// 		}
// 		i++
// 	}
// 	i, j := 0, 0
// 	for i < ns {
// 		if nums[i] != -111 {
// 			nums[j] = nums[i]
// 			j++
// 		}
// 		i++
// 	}

// 	for i := 0; i < ns; i++ {
// 		fmt.Printf("%d ", nums[i])
// 	}
// 	return ns - cnt
// }

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(nums)

}
