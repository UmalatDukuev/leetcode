package main

import (
	"fmt"
)

func spacing(str []string, maxWidth int, wordsLen int) string {
	maxSpace := maxWidth - wordsLen
	fmt.Println("ssssss: ", maxSpace)
	return ""
}

func fullJustify(words []string, maxWidth int) []string {
	spaces := make([]int, 1)
	wordsLen := len(words)
	strLen := 0

	for i := 0; i < wordsLen; i++ {
		strLen += len(words[i]) + 1
		if strLen > maxWidth+1 {
			strLen = 0
			spaces = append(spaces, i)
			i--
		}
	}
	spaces = append(spaces, wordsLen)

	//result_strings := make([]string, len(spaces))
	//substr := ""

	for i := 0; i < len(spaces)-1; i++ {
		subStrs := make([]string, 0)
		wordsLen := 0
		j := int(spaces[i])
		for j < int(spaces[i+1]) {
			wordsLen += len(words[j])
			subStrs = append(subStrs, words[j])
			//substr += words[j] + " "
			j++
		}
		fmt.Println(subStrs)
		spacing(subStrs, maxWidth, wordsLen)

		//substr += words[j]
		//fmt.Println(substr)
		//substr = "\"" + substr + "\""
		//result_strings[i] = substr
		//substr = ""
	}

	//fmt.Println(result_strings)

	return []string{}
}

func main() {
	words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth := 20
	fullJustify(words, maxWidth)
}
