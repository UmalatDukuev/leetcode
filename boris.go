package main

import (
	"fmt"
	"strconv"
)

func checkPowTwo(num int) (bool, int) {
	strin := ""
	for num > 0 {
		strin += strconv.Itoa(num % 2)
		num /= 2
	}
	strLen := len(strin) - 1
	for i := 0; i < strLen; i++ {
		if i != strLen && strin[i] == '1' {
			return false, -1
		}
	}
	return true, strLen
}

func findThree(num int) int {
	var mas []int
	cnt := 0
	for i := num; i > 0; i-- {
		flag, _ := checkPowTwo(i)
		if flag == true && len(mas) < 3 && cnt+i <= num {
			cnt += i
			mas = append(mas, i)
		}
	}
	if len(mas) != 3 {
		return -1
	}
	return cnt
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		fmt.Println(findThree(a))
	}
}
