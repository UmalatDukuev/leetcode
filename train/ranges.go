package main

import (
	"fmt"
	"sort"
)

func ranges(min, max int, busy []int) [][]int {

	sort.Ints(busy)
	fmt.Println(busy)
	result := make([][]int, 1)
	busyInRange := make([]int, 0)

	for _, rang := range busy {
		if rang >= min && rang <= max {
			busyInRange = append(busyInRange, rang)
		}
	}
	result[0] = append(result[0], min)

	for _, rang := range busyInRange {
		fmt.Println(rang)

	}
	return result

}

func main() {
	min := 30000
	max := 32000
	busy := []int{31000, 32000, 20000}
	fmt.Println(ranges(min, max, busy))
}
