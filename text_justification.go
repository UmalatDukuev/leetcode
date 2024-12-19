package main

import (
	"fmt"
)

func fullJustify(words []string, maxWidth int) []string {
	spaces := make([]int, 1)
	wordsLen := len(words)
	strLen := 0
	cnt := 0
	for _, val := range words {
		cnt += len(val)
		fmt.Printf("%s: %d\n", val, len(val))
		//fmt.Printf("Total %d\n", cnt)

	}
	fmt.Println("|||||||||||||||||||||||||||||||||||||||")

	for i := 0; i < wordsLen; i++ {
		strLen += len(words[i]) + 1
		cnt += 1
		if strLen > maxWidth+1 {
			fmt.Println(strLen)
			fmt.Println(words[i])

			cnt = 0
			strLen = 0
			spaces = append(spaces, i)
			i--
		}
	}
	fmt.Println(spaces)
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
