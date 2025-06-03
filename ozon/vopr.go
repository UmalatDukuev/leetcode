package main

import "fmt"

func main() {
	// slice := []int{1, 2}
	slice := make([]int, 1, 2)
	fmt.Println(slice)
	append1(*slice, 11)
	fmt.Println(slice)
	slice = append(slice, 11)
	fmt.Println(slice)

	// mutate(slice, 1, 3)
	// fmt.Println(slice)

}

func append1(sl *[]int, val int) {
	sl = append(sl, val)
}

func mutate(sl []int, idx int, val int) {
	sl[idx] = val
}
