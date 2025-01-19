package main

import (
	"bufio"
	"fmt"
	"os"
)

func deletePosition(num string, pos int) string {
	num = num[:pos] + num[pos+1:]
	return num
}

func findMax(num string) string {
	n := num
	str_len := len(n)
	if str_len == 1 {
		return "0"
	}
	ind := 0
	i := 0
	for i < str_len-1 {
		a := string(n[i])
		b := string(n[i+1])
		if a < b {
			ind = i
			return deletePosition(n, ind)
		}
		i++
	}
	return deletePosition(n, str_len-1)
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var a int
	fmt.Fscan(in, &a)
	numbers := make([]string, a)
	for i := 0; i < a; i++ {
		fmt.Fscan(in, &numbers[i])
	}
	for i, val := range numbers {
		numbers[i] = findMax(val)
	}
	for _, val := range numbers {
		fmt.Println(val)
	}
}
