package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	k := m + n - 1
	m_ptr := m - 1
	n_ptr := n - 1
	if m == 0 {
		for i := 0; i < len(nums2); i++ {
			nums1[i] = nums2[i]
		}
	} else if n == 0 {
		return
	} else {
		for m_ptr >= 0 && n_ptr >= 0 {
			if nums1[m_ptr] > nums2[n_ptr] {
				nums1[k] = nums1[m_ptr]
				m_ptr--
			} else {
				nums1[k] = nums2[n_ptr]
				n_ptr--
			}
			k--
		}
	}
	for n_ptr >= 0 {
		nums1[k] = nums2[n_ptr]
		k--
		n_ptr--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	for i := 0; i < len(nums1); i++ {
		fmt.Printf("%d ", nums1[i])
	}
}
