package main

import (
	"fmt"
	"sort"
)

func ranges(min, max int, busy []int) [][]int {

	sort.Ints(busy)
	result := make([][]int, 0)
	busyInRange := make([]int, 0)

	for _, rang := range busy {
		if rang >= min && rang <= max {
			busyInRange = append(busyInRange, rang)
		}
	}
	currentMin := min
	for _, b := range busyInRange {
		if b > currentMin {
			result = append(result, []int{currentMin, b - 1})

		}
		currentMin = b + 1
	}
	if currentMin <= max {
		result = append(result, []int{currentMin, max})
	}

	return result

}

func main() {
	min := 30000
	max := 32000
	busy := []int{30100, 30200, 20000}
	fmt.Println(ranges(min, max, busy))
}
