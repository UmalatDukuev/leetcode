package main

import "fmt"

func majorityElement(nums []int) int {
	curr, cnt := 0, 0

	for _, num := range nums {
		if cnt == 0 {
			curr = num
			cnt++
		} else if curr == num {
			cnt++
		} else {
			cnt--
		}
	}
	return curr
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(majorityElement(nums))
}
