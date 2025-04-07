package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkString(strg string) string {
	var flag string
	strLen := len(strg)
	flag = "YES"

	if strLen == 1 {
		return flag
	}

	first := strg[0]
	for i := 0; i < len(strg); i++ {
		if i%2 == 0 && strg[i] != first {
			flag = "NO"
		}
	}
	return flag
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var strg string
		fmt.Fscanln(reader, &strg)
		result := checkString(strg)
		fmt.Fprintln(writer, result)
	}
}
