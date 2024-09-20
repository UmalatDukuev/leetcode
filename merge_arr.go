package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	k := &nums1[m+n-1]
	m_ptr := &nums1[m+n-1]
	n_ptr := &nums2[m-1]

	for i := m + m - 1; i > 0; i-- {
		if n_ptr <  
	}
	fmt.Println(*m_ptr)
	fmt.Println(*n_ptr)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
}
