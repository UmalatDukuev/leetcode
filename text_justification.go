package main

import (
	"fmt"
)

func fullJustify(words []string, maxWidth int) []string {
	spaces := make([]int, 1)
	strLen := 0
	cnt := 0
	for ind, val := range words {
		strLen += len(val) + 1
		cnt += 1
		if strLen > maxWidth {
			fmt.Println(strLen)
			fmt.Println(val)

			cnt = 0
			strLen = 0
			spaces = append(spaces, ind)
		}
	}
	// st := 0
	// for st < maxWidth {
	// 	st += len(val)
	// }
	// strings := make([]string, 0)
	// substr := ""
	// fmt.Println(spaces)
	// for ind, _ := range spaces {
	// 	fmt.Println(ind)
	// 	for i := int(spaces[ind]); i < int(spaces[ind+1])-1; i++ {
	// 		substr += words[i]
	// 	}
	// 	strings = append(strings, substr)
	// }
	// fmt.Println(strings)

	return []string{}
}

func main() {
	words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth := 20
	fmt.Println(fullJustify(words, maxWidth))
}
