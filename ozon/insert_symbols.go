package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkString(strg string) string {
	strLen := len(strg)

	if strLen == 1 {
		return "YES"
	}
	if strLen == 2 {
		if strg[0] == strg[1] {
			return "YES"
		} else {
			return "NO"
		}
	}

	first := strg[0]
	last := strg[strLen-1]

	if last != first {
		return "NO"
	}
	for i := 1; i < len(strg)-1; i++ {
		if strg[i] != first {
			if strg[i-1] != first || strg[i+1] != first {
				return "NO"
			}
		}
	}
	return "YES"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)
	//t = 1
	for i := 0; i < t; i++ {
		var strg string
		fmt.Fscanln(reader, &strg)
		result := checkString(strg)
		fmt.Fprintln(writer, result)
	}
}
