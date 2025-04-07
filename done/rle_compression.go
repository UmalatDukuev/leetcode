package main

import (
	"strconv"
	"strings"
)

func compress_string(strg string) string {
	var sb strings.Builder
	count := 1
	sb.WriteString(string(strg[0]))
	for i := 1; i < len(strg); i++ {
		if strg[i-1] != strg[i] {
			sb.WriteString(strconv.Itoa(count))
			sb.WriteString(string(strg[i]))
			count = 1
		} else {
			count++
		}
	}
	sb.WriteString(strconv.Itoa(count))
	return sb.String()
}

func main() {
	strg := "AABBBBCDD"
	println(compress_string(strg))
}
