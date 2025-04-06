package main

import (
	"strconv"
	"strings"
)

func decode_text(strg string) string {
	var sb strings.Builder
	for i := 0; i < len(strg); i++ {
		if strg[i] >= '1' && strg[i] <= '9' {
			num, _ := strconv.Atoi(string(strg[i]))
			println(num)
			// if rune(strg[i]) == '[' {
			// 	for rune(strg[i]) != ']' {
			// 		sb.WriteString(string(strg[i]))
			// 	}
			// }
		}
	}
	return sb.String()
}

func main() {
	text := "3[a]2[bc]"
	println(decode_text(text))
}
