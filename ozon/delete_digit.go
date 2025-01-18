package main

import (
	"fmt"
	"strconv"
)

func deleteMinInNumber(num int) int {
	if num/10 == 0 {
		return 0
	}
	min := 100
	k := 0
	n := num
	for i := 0; num > 0; i++ {
		p := num % 10
		if p <= min {
			min = p
			k = i
		}
		num /= 10
	}
	st := strconv.Itoa(n)
	le := len(st)
	st = st[:(le-k-1)] + st[(le-k):]
	n, _ = strconv.Atoi(st)
	return n
}

func main() {
	// var in *bufio.Reader
	// var out *bufio.Writer
	// in = bufio.NewReader(os.Stdin)
	// out = bufio.NewWriter(os.Stdout)
	// defer out.Flush()
	// var a int
	// fmt.Fscan(in, &a)
	// numbers := make([]int, a)
	// for i := 0; i < a; i++ {
	// 	fmt.Fscan(in, &numbers[i])
	// }
	numbers := []int{987241, 3233, 3432}
	// max := -1
	for i, val := range numbers {
		numbers[i] = deleteMinInNumber(val)
		// if numbers[i] > max {
		// 	max = numbers[i]
		// }
	}
	for _, val := range numbers {
		fmt.Println(val)
	}

}
