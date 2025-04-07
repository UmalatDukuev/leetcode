package main

import (
	"fmt"
)

func intersectionWithRepetition(a, b []int) []int {
	counter := make(map[int]int)
	for _, elem := range a {
		if _, ok := counter[elem]; !ok {
			counter[elem] = 1
		} else {
			counter[elem] += 1
		}
	}
	var result []int
	for _, elem := range b {
		if count, ok := counter[elem]; ok && count > 0 {
			counter[elem]--
			result = append(result, elem)
		}
	}

	return result
}

func intersectionWithoutRepetition(a, b []int) []int {
	counter := make(map[int]struct{})
	for _, elem := range a {
		counter[elem] = struct{}{}
	}
	var result []int
	for _, elem := range b {
		if _, ok := counter[elem]; ok {
			result = append(result, elem)
			delete(counter, elem)
		}
	}
	return result
}

func main() {
	arr1 := []int{1, 3, 5, 4, 4, 2}
	arr2 := []int{4, 5, 6, 4}
	result := intersectionWithRepetition(arr1, arr2)
	fmt.Println(result)
	arr1 = []int{44, 6, 5, 1, 4, 2, 2, 33}
	arr2 = []int{9, 0, 5, 7, 6, 44}
	result = intersectionWithoutRepetition(arr1, arr2)
	fmt.Println(result)
}
